package db_mod

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhashkevych/todo-app/modules/user"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



/*структра*/
type db struct {

	collection *mongo.Collection//Наша коллекция
	logger     loging.Logger//логер

}



/*метод создания юзера*/
func (d *db) Create(ctx context.Context, user user.User_User) (string, error) {

	d.logger.Debug("create user")//логируем
	result_insert_one, err := d.collection.InsertOne(ctx, user)

	if err != nil {

		return "", fmt.Errorf("failed to create user due to error: %v", err)

	}

	d.logger.Debug("convert InsertedID to ObjectID")//логируем

	oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)

	/*проверка если все окей ибо без него он возвращает */
	if ok {

		return oid.Hex(), nil//возвращаем то что нужно вернуть

	}


	d.logger.Trace(user)//логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)

}

func (d *db) FindAll(ctx context.Context) (user_user []user.User_User, err error) {

	cursor, err := d.collection.Find(ctx, bson.M{})

	if cursor.Err() != nil {

		return user_user, fmt.Errorf("failed to find all users due to error: %v", err)

	}

	if err = cursor.All(ctx, &user_user); err != nil {

		return user_user, fmt.Errorf("failed to read all documents from cursor. error: %v", err)

	}

	return user_user, nil

}


func (d *db) FindOneUser(ctx context.Context, id string) (user_user User_User, err error) {



	id_object, err := primitive.ObjectIDFromHex(id)//получаем либо обджект Id либо ошибку

	if err != nil {

		return user_user, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)

	}

	filter := bson.M{"_id": id_object}

	result := d.collection.FindOne(ctx, filter)//коллекция FindOne

	if result.Err() != nil {

		if errors.Is(result.Err(), mongo.ErrNoDocuments) {

			return user_user, nil

		}
		return user_user, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)

	}
	if err = result.Decode(&user_user); err != nil {

		return user_user, fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)

	}

	/* select по преглашениям юзеров*/
	invitations := bson.D{
		{user_user.UserInvitations,result},//select UserInvitations

	}

	select_invitations, err := d.collection.Find(

		context.Background(),

		bson.D{

			{user_user.ID ,result},//where ID= id юзера

		},

		options.Find().SetProjection(invitations),

	)

	d.logger.Trace(select_invitations)


	return user_user, nil//если все нормально возвращаем user_user

}

/*метод добавления юзера*/
func (d *db) Update(ctx context.Context, user user.User_User) error {



	objectID, err := primitive.ObjectIDFromHex(user.ID)//id в Hex

	/*Eсли не удалось преобразовать идентификатор пользователя в ObjectID*/
	if err != nil {

		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", user.ID)

	}

	filter := bson.M{"_id": objectID}// зоздаем фильтр

	userBytes, err := bson.Marshal(user)//маршалим юзера

	/*если неудолось заморшалить юзера*/
	if err != nil {

		return fmt.Errorf("failed to marhsal user. error: %v", err)

	}

	var updateUserObj bson.M//превращаем юзер байты в обькект бейчона

	err = bson.Unmarshal(userBytes, &updateUserObj)//превращаем юзер байты в обькект бейчона

	if err != nil {

		return fmt.Errorf("failed to unmarshal user bytes. error: %v", err)

	}

	delete(updateUserObj, "_id")//удаляем поле чтобы оно смогло обновиться в базе

	/*создать квери на обновление */
	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)//резельтат UpdateOne

	if err != nil {

		return fmt.Errorf("failed to execute update user query. error: %v", err)

	}
	/*смотрим нашли ли что нибудь если нет возвращаем nil */
	if result.MatchedCount == 0 {
		return nil
	}

	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil

}

/* метод удаления*/
func (d *db) Delete(ctx context.Context, id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)//id в Hex

	/*Eсли не удалось преобразовать идентификатор пользователя в ObjectID*/
	if err != nil {

		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", id)

	}

	filter := bson.M{"_id": objectID}//фильтер через bson

	result, err := d.collection.DeleteOne(ctx, filter)//резельтат DeleteOne

	/* если не удалось выполнить запрос*/
	if err != nil {

		return fmt.Errorf("failed to execute query. error: %v", err)

	}

	/*смотрим нашли ли что нибудь если нет возвращаем nil */
	if result.DeletedCount == 0 {
		return nil
	}

	d.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil//возвращаем nil

}

/*конструктор */
func NewStorage(database *mongo.Database, collection string, logger loging.Logger) user.Storage {

	return &db{

		collection: database.Collection(collection),//бд

		logger:     logger,

	}

}
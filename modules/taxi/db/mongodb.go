package db



import (
	"context"
	"fmt"
	"github.com/zhashkevych/todo-app/modules/taxi"
	"github.com/zhashkevych/todo-app/modules/user"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

/*структра*/
type db struct {

	collection *mongo.Collection//Наша коллекция
	logger     loging.Logger//логер

}

func (d db) Payment(ctx context.Context, db taxi.Cars, u user.User_User ) (string, error) {

	bonus:=u.Bonus//бонусы

	/*вычетаем 1000 бонусов*/
	subtract_bonus := bonus - 1000

	subtract_bonus_xt := reflect.TypeOf(subtract_bonus).Kind()

	/*смотрим достаточно ли нам бонусов проверяем тип чтобы число бонусов не было меньше нуля*/
	if reflect.Uint64 != subtract_bonus_xt {
		panic(subtract_bonus)
	}
	if subtract_bonus ==subtract_bonus {

		d.logger.Debug("create ") //логируем
		result_insert_one, err := d.collection.InsertOne(ctx,taxi.Cars{Driver: 1,Passenger: 1})

		if err != nil {

			return "", fmt.Errorf("failed to create Taxi due to error: %v", err)
		}

		d.logger.Debug("convert InsertedID to ObjectID")//логируем

		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает */
		if ok {

			return oid.Hex(), nil//возвращаем то что нужно вернуть

		}

	}

	d.logger.Trace(db)//логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", db)

}

func (d db) Create(ctx context.Context, db taxi.Cars) (string, error) {
	d.logger.Debug("create taxi")//логируем
	result_insert_one, err := d.collection.InsertOne(ctx, db)

	if err != nil {

		return "", fmt.Errorf("failed to create Taxi due to error: %v", err)

	}

	d.logger.Debug("convert InsertedID to ObjectID")//логируем
	oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
	/*проверка если все окей ибо без него он возвращает */
	if ok {

		return oid.Hex(), nil//возвращаем то что нужно вернуть

	}


	d.logger.Trace(db)//логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d db) Update(ctx context.Context, db taxi.Cars) error {

	objectID, err := primitive.ObjectIDFromHex(db.ID)//id в Hex

	/*Eсли не удалось преобразовать идентификатор пользователя в ObjectID*/
	if err != nil {

		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", db.ID)

	}

	filter := bson.M{"_id": objectID}// зоздаем фильтр

	userBytes, err := bson.Marshal(db)//маршалим taxi

	/*если неудолось заморшалить taxi*/
	if err != nil {

		return fmt.Errorf("failed to marhsal user. error: %v", err)

	}

	var updateUserObj bson.M//превращаем taxi байты в обькект бейчона

	err = bson.Unmarshal(userBytes, &updateUserObj)//превращаем taxi байты в обькект бейчона

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

func (d db) Delete(ctx context.Context, id string) error {

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
func NewStorage(database *mongo.Database, collection string, logger loging.Logger) taxi.TaxiDB {

	return &db{

		collection: database.Collection(collection),//бд

		logger:     logger,

	}

}
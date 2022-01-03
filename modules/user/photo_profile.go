package user

import (
	"context"
	"fmt"
	_ "github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
)

func (d *db) CreatePaidPhoto(ctx context.Context, bonus User_User) (string, error) {


	d.logger.Debug("create ")//логируем


	if bonus.ProfilePhotos!=nil {

		profile_bonus, err:= d.Collection.InsertOne(ctx, bonus.Bonus) //добавление бонусов за фото

		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)
		}

		d.logger.Debug("convert InsertObjectID") //логируем

		oid, ok := profile_bonus.InsertedID.(primitive.ObjectID)

		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть
		}

	}

	d.logger.Trace(bonus.Bonus) //логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: ,,kk")
}


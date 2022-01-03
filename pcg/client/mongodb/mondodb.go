package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {

	var mongoDBURL string//Переменная с авторизацией
	var isAuth bool//переменная для проверки есть ли регистрация в бд

	/**
	*проверка есть ли авторизация или ее нет
	*/
	if username == "" && password == "" {
		/*если ли нет авторизации */
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)

	} else {
		/*если ли есть авторизация */
		isAuth = true

		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)//передаем данные в базу

	if isAuth {

		if authDB == "" {
			//присваиваем authDB = database ибо authDB не может быть пустой
			authDB = database

		}

		clientOptions.SetAuth(options.Credential{

			AuthSource: authDB,
			Username:   username,
			Password:   password,

		})

	}
	//сойдинение с бд
	client, err := mongo.Connect(ctx, clientOptions)

	/*проверки на ошибка*/
	if err != nil {

		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)

	}

	if err = client.Ping(ctx, nil); err != nil {

		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)

	}

	return client.Database(database), nil//возвращаем бд

}
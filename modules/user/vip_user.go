package user

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


/*создания оплаты на сайте*/
func Vip() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(400),//400 центов или же 4 доллара

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте*/
func VipSuper() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(1000),//1000 центов или же 10 долларов

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте*/
func VipMega() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(2000),//2000 центов или же 20 долларов

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}

func Anon()  *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(9900),//9900 центов или же 99 долларов

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*метод создания бонуса*/
func (d *db) CreatVip(ctx context.Context,bonus User_User) (string, error) {
	if Vip()!= nil {
		d.logger.Debug("create ")
		//логируем
		result_insert_one_Frame, err := d.Collection.InsertOne(ctx, User_User{Frame: "100"})
		result_insert_one_Bonus, err := d.Collection.InsertOne(ctx, User_User{Bonus: 1000}) //300 бонусов
		/*Сделать чтобы цвет в чате менялся*/
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one_Bonus.InsertedID.(primitive.ObjectID)
		oi, ok := result_insert_one_Frame.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {
			return oi.Hex(),nil
			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}

	}
	if VipSuper() != nil {

		d.logger.Debug("create ")
		//логируем
		result_insert_one_Username, err := d.Collection.InsertOne(ctx, User_User{Username: "Gold"})
		result_insert_one_Frame, err := d.Collection.InsertOne(ctx, User_User{Frame: "100"})
		result_insert_one_Bonus, err := d.Collection.InsertOne(ctx, User_User{Bonus: 2000}) //300 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one_Bonus.InsertedID.(primitive.ObjectID)
		oi, ok := result_insert_one_Frame.InsertedID.(primitive.ObjectID)
		i, ok := result_insert_one_Username.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {
			return oi.Hex(),nil
			return oid.Hex(), nil //возвращаем то что нужно вернуть
			return i.Hex(),nil

		}

	}
	if VipMega() !=nil{
		d.logger.Debug("create ")
		//логируем
		resultInsertOneUsername, err := d.Collection.InsertOne(ctx, User_User{Username: "Gold"})
		resultInsertOneFrame, err := d.Collection.InsertOne(ctx, User_User{Frame: "100"})
		resultInsertOneBonus, err := d.Collection.InsertOne(ctx, User_User{Bonus: 2000}) //300 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := resultInsertOneBonus.InsertedID.(primitive.ObjectID)
		oi, ok := resultInsertOneFrame.InsertedID.(primitive.ObjectID)
		i, ok := resultInsertOneUsername.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {
			return i.Hex(),nil
			return oi.Hex(),nil
			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}

	}
	if Anon() !=nil{
		d.logger.Debug("create ")
		//логируем
		resultInsertOneUsername, err := d.Collection.InsertOne(ctx, User_User{Username: "Gold"})
		resultInsertOneFrame, err := d.Collection.InsertOne(ctx, User_User{Frame: "100"})
		resultInsertOneBonus, err := d.Collection.InsertOne(ctx, User_User{Bonus: 2000}) //300 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := resultInsertOneBonus.InsertedID.(primitive.ObjectID)
		oi, ok := resultInsertOneFrame.InsertedID.(primitive.ObjectID)
		i, ok := resultInsertOneUsername.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {
			return i.Hex(),nil
			return oi.Hex(),nil
			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}

	}

	d.logger.Trace(bonus.Bonus) //логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: ,,kk")

}



package user

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "os/user"
)

type db struct {

	Collection *mongo.Collection //Наша коллекция
	logger     loging.Logger     //логер

}


/*ИЗВИНЯЮСЬ ЗА ПЛОХОЙ КОД НЕ БЫЛО ВРЕМИНИ РАЗБИРАТЬСЯ С ОПЛАТОЙ stripe когда у вас будет время разберитесь
ибо это весьма констыльное решаение =( Я просто сделала разные сиансы оплаты для разного колличества бонусов
если найдете как сделать проще сделайте
 */


/*создания оплаты на сайте  добавление для 300 монет*/
func CreateBonus300Stripe() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(100),//100 центов или же один доллар

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте  добавление для 1000 монет*/
func CreateBonus1000Stripe() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(200),//200 центов или же 2 доллара

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте добавление для 2000 монет */
func CreateBonus2000Stripe() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(300),//300 центов или же 3 доллара

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте*/
func CreateBonus4000Stripe() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{


		Amount:   stripe.Int64(500),//500 центов или же 5 долларов

		Currency: stripe.String(string(stripe.CurrencyUSD)),//валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",

		}),

		ReceiptEmail: stripe.String("your@mail.com"),//пока такой емиал

	}

	return params_sp //возвращаем оплату


}
/*создания оплаты на сайте*/
func CreateBonus10000Stripe() *stripe.PaymentIntentParams {

	/*пораметры оплаты*/
	params_sp := &stripe.PaymentIntentParams{

		Amount: stripe.Int64(1000), //1000 центов или же 10 долларов

		Currency: stripe.String(string(stripe.CurrencyUSD)), //валюта доллары

		PaymentMethodTypes: stripe.StringSlice([]string{

			"card",
		}),

		ReceiptEmail: stripe.String("your@mail.com"), //пока такой емиал

	}

	return params_sp //возвращаем оплату

}

/*метод создания бонуса*/
func (d *db) CreateBonus(ctx context.Context,bonus User_User) (string, error) {

	if CreateBonus300Stripe() != nil {

		d.logger.Debug("create ")                                                    //логируем
		result_insert_one, err := d.Collection.InsertOne(ctx, User_User{Bonus: 300}) //300 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}
	}

	if CreateBonus1000Stripe() !=nil {
		d.logger.Debug("create ")                                                     //логируем
		result_insert_one, err := d.Collection.InsertOne(ctx, User_User{Bonus: 1000}) //1000 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}

	}

	if CreateBonus2000Stripe() !=nil {
		d.logger.Debug("create ")                                                     //логируем
		result_insert_one, err := d.Collection.InsertOne(ctx, User_User{Bonus: 2000}) //2000 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}

	}

	if CreateBonus4000Stripe() !=nil {
		d.logger.Debug("create ")                                                     //логируем
		result_insert_one, err := d.Collection.InsertOne(ctx, User_User{Bonus: 4000}) //4000 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}
	}

	if CreateBonus10000Stripe() !=nil {
		d.logger.Debug("create ")                                                     //логируем
		result_insert_one, err := d.Collection.InsertOne(ctx, User_User{Bonus: 2000}) //10000 бонусов
		if err != nil {

			return "", fmt.Errorf("failed to create user due to error: %v", err)

		}
		d.logger.Debug("convert InsertedID to ObjectID") //логируем
		oid, ok := result_insert_one.InsertedID.(primitive.ObjectID)
		/*проверка если все окей ибо без него он возвращает*/
		if ok {

			return oid.Hex(), nil //возвращаем то что нужно вернуть

		}
	}

	d.logger.Trace(bonus.Bonus) //логируем
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: ,,kk")

}
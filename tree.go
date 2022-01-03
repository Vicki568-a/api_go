package main

import (
	"context"
	"github.com/zhashkevych/todo-app/modules/config"
	"github.com/zhashkevych/todo-app/modules/user"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/zhashkevych/todo-app/modules/user/db_mod"
	"github.com/zhashkevych/todo-app/pcg/client/mongodb"
	"github.com/zhashkevych/todo-app/pcg/loging"
	_ "github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson"

	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type db struct {

	Collection *mongo.Collection //Наша коллекция
	logger     loging.Logger     //логер

}

/*Добавления транзакции бонусов к приглашениям*/

func  (d db)addTransaction(ctx context.Context, bonus user.User_User)  {

	start_percent := 10 //стартовое число отсчета процентов

	step := 2  //Число, на которое будет делится $current_percent

	current_percent := start_percent //Текущий процент

	total_expense := 0 //Общее кол-во розданных бонусов

	log.Printf(string(total_expense))


	/*Обход пригласивщих пользователей и раздача бонусов*/

	for  invitations:= range bonus.UserInvitations {

		log.Printf(string(invitations))

		/*Процент текущего пользователя*/

		invitations_percent := current_percent / step

		current_percent = invitations_percent

		/*Бонус текущего пользователя*/

		invitations_bonus := bonus.Bonus / 100 * invitations_percent

		/*Сохранение результата*/

		logger := loging.GetLogger() //получаем логер

		cfg := config.GetConfig() //получаем конфиг

		cfgMongo := cfg.Mongodb

		mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
			cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)

		storage :=db_mod.NewStorage(mongoDBClient, cfg.Mongodb.Collection, logger)

		//result, err := d.collection.UpdateOne(ctx, filter, update)//резельтат UpdateOne

		new_user_id, err := storage.FindOneUser(context.Background(), bonus.ID)


		if err != nil {

			panic(err)

		}


		/* select по преглашениям юзеров
		SELECT bonus FROM users WHERE id = invitations_id
		*/
		calculeted_bonus := bson.D{
			{string(bonus.Bonus), 1}, //select UserInvitations

		}

		select_calculeted_bonus, err := d.Collection.Find(

			context.Background(),

			bson.D{

				{bonus.ID, bonus.UserInvitations}, //where ID= id юзера

			},

			options.Find().SetProjection(calculeted_bonus),

		)

		if err != nil {

			panic(err)

		}

		/*UPDATE `users` SET `bonus` = select_calculeted_bonus WHERE id = invitations_id*/
		filter := d.Collection.FindOne(ctx,

			bson.D{

			{string(bonus.Bonus), select_calculeted_bonus},

			{bonus.ID, bonus.UserInvitations},

			},

		)

		result := d.Collection.FindOne(ctx, filter) //коллекция FindOne

		logger.Info(result)                   //логируем
		logger.Trace(select_calculeted_bonus) //Логируем


		/* select по преглашениям юзеров
		SELECT bonus FROM users WHERE id = new_user_id
		*/

		calculeted_bonus_select := bson.D{
			{string(bonus.Bonus), 1}, //select UserInvitations

		}

		select_calculeted_bonus_one, err := d.Collection.Find(

			context.Background(),

			bson.D{

				{bonus.ID, new_user_id}, //where ID= id юзера

			},


			options.Find().SetProjection(calculeted_bonus_select),

		)
		logger.Info(select_calculeted_bonus_one)
		bonus_bonus:=0
		bonus_bonus +=bonus.Bonus-invitations_bonus

		calculeted_bonus_user := bson.D{
			{string(bonus.Bonus), 1}, //select UserInvitations

		}
		logger.Info(calculeted_bonus_user)
		select_calculeted_bonus_user:= d.Collection.FindOne(ctx,

			bson.D{

				{string(bonus.Bonus), bonus_bonus},

				{bonus.ID, new_user_id},

			},

		)
		logger.Info(select_calculeted_bonus_user)

		result = d.Collection.FindOne(ctx, filter) //коллекция FindOne

		logger.Info(result)

	}




}


/*



func (d db_mod)addTransaction(ctx context.Context, bonus User_User) (int, error) {


	invitations_id := d.getUserInvitations
	projection := bson.D{
		{"bonus", 1},
	}
	calculeted_bonus, err := d.collection.Find(
		
		context.Background(),
		
		bson.D{
			
			{"id", invitations_id},
		},
		
		options.Find().SetProjection(projection),
	
		)
	
	report := d.calculateBonus

	calculeted_bonus =invitations_id['follower_bonus']
	
$calculeted_bonus= (float)mysqli_fetch_array(mysqli_query($conn, "SELECT bonus FROM users WHERE id = $user_id LIMIT 1"))['bonus'];
$calculeted_bonus+=$bonus-(float)$report[ "total_expense" ];


mysqli_query($conn, "UPDATE `users` SET `bonus` = $calculeted_bonus WHERE id = $user_id");


   followers_id:=getFollowers($user_id);

   $report = calculateBonus($bonus, $followers_id);

   foreach ($report[ "followers" ] as $follower){
   $follower_id=$follower['follower_id'];


	// 1. Проходишься по всем report[ "followers" ], и начисляешь им бонусы (можешь параллельно вести лог)
	// 2. Вычитаешь бонусы у текущего пользователя ($user_id); кол-во бонусов берешь у report[ "total_expense" ]
}
*/

package advert

import "github.com/zhashkevych/todo-app/modules/user"

/*таблицы обьявлений*/
type Advert struct {
	Username		string `json:"username" bson:"username"`
	ID				string `json:"id" bson:"_id,omitempty"`
	Advertname		string `json:"advert_name" bson:"advert_name"`
	Place			string `json:"place" bson:"place"`
	Time			string `json:"time" bson:"time"`
	Date			string `json:"date" bson:"date"`
	Cite			string `json:"cite" bson:"cite"`
	Description		string `json:"description" bson:"description"`
	Photo			string `json:"photo" bson:"photo"`
	PhotoPlace		string `json:"photo_place" bson:"photo_place"`
	Like			string `json:"like" bson:"like"`
	UserID			user.User_User`json:"UserID" bson:"UserID"`

}

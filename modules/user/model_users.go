package user

import (
	"image"
	"time"
)

type User_User struct {
	ID				string `json:"id" bson:"_id,omitempty"`
	Email			string `json:"email" bson:"email"`
	Username		string `json:"username" bson:"username"`
	ProfileAvatar	image.Image `json:"avatar" bson:"avatar"`
	ProfilePhotos	image.Image `json:"photo" bson:"photo"`
	PasswordHash	int `json:"-" bson:"password"`
	Bonus			int `json:"bonus" bson:""`
	UserInvitations	string `json:"user_invitations" bson:""`
	Frame			string `json:"frame" bson:"frame"`
	Vip				string `json:"vip" bson:"vip"`
	Name			string `json:"Name" bson:"Name"`
	Surnames		string `json:"Surnames" bson:"Surnames"`
	Surname			string `json:"Surname" bson:"Surname"`
	DateOfBirth		string `json:"DateOfBirth" bson:"DateOfBirth"`
	Location		string `json:"Location" bson:"Location"`
	Floor			string `json:"Floor" bson:"Floor"`
	BodyType		string `json:"BodyType" bson:"BodyType"`
	Growth			string `json:"Growth" bson:"Growth"`
	EyeColor		string `json:"EyeColor" bson:"EyeColor"`
	Orientation		string `json:"Orientation" bson:"Orientation"`
	Religion		string `json:"Religion" bson:"Religion"`
	Hobby			string `json:"Hobby" bson:"Hobby"`
	ZodiacSign		string `json:"ZodiacSign" bson:"ZodiacSign"`
	FoodPreferences	string `json:"FoodPreferences" bson:"FoodPreferences"`
	Pet				string `json:"Pet" bson:"Pet"`
	HaveChildren	string `json:"HaveChildren" bson:"HaveChildren"`
	Education		string `json:"Education" bson:"Education"`
	Works			string `json:"Works" bson:"Works"`
	Wage			string `json:"Wage" bson:"Wage"`


}

func (u User_User) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (u User_User) Done() <-chan struct{} {
	panic("implement me")
}

func (u User_User) Err() error {
	panic("implement me")
}

func (u User_User) Value(key interface{}) interface{} {
	panic("implement me")
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
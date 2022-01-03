package user

import (
	"context"
)
/*интерфейс*/
type Storage interface {
	Create(ctx context.Context, user User_User) (string, error)
	FindAll(ctx context.Context) (u []User_User, err error)
	FindOneUser(ctx context.Context, id string) (User_User, error)
	Update(ctx context.Context, user User_User) error
	Delete(ctx context.Context, id string) error

}

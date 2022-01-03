package taxi

import (
	"context"
	"github.com/zhashkevych/todo-app/modules/user"
)

/*интерфейс*/
type TaxiDB interface {

	Create(ctx context.Context,db Cars) (string, error)
	Update(ctx context.Context, db Cars) error
	Delete(ctx context.Context, id string) error
	Payment(ctx  context.Context, db Cars, u user.User_User) (string, error)

}
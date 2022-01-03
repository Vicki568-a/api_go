package advert

import (
	"golang.org/x/net/context"
)
/*интерфейс*/
type Storage interface {
	Create(ctx context.Context, user Advert) (string, error)
	FindAll(ctx context.Context) (u []Advert, err error)
	FindOne(ctx context.Context, id string) (Advert, error)
	Update(ctx context.Context, user Advert) error
	Delete(ctx context.Context, id string) error

}
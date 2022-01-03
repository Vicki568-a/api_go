package advert

import (
	"context"
	"github.com/zhashkevych/todo-app/pcg/loging"
)

type ServiceAdver struct {
	storage Storage
	logger loging.Logger
}

func (s *ServiceAdver) CreateAdvert(ctx context.Context, dto Advert) (u Advert, err error) {
	// TODO for next one
	return
}
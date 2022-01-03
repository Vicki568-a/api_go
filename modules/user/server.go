package user

import (
	"context"
	"github.com/zhashkevych/todo-app/pcg/loging"
)

type Service struct {
	storage Storage//Для работы с базой
	logger loging.Logger
}
/*метод создание*/
func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User_User, err error) {
	// TODO for next one
	return
}
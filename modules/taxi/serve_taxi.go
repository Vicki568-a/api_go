package taxi


import (
	"context"
"github.com/zhashkevych/todo-app/pcg/loging"
)

type Service_Taxi struct {
	storage TaxiDB//Для работы с базой
	logger loging.Logger
}
/*метод создание*/
func (s *Service_Taxi) CreateTaxi (ctx context.Context, dto TaxiDB) (u TaxiDB, err error) {
	// TODO for next one
	return
}
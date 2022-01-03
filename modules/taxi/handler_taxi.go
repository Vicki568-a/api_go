package taxi


import (
	"github.com/julienschmidt/httprouter"
	"github.com/zhashkevych/todo-app/modules/hendler"
	_ "github.com/zhashkevych/todo-app/modules/hendler"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"net/http"
)

var _ hendler.Handler = &handler{} //подсказка реализует ли структура интерфейс с помощью Register

const (

)

type handler struct {
	logger loging.Logger
}

func NewHandler(logger loging.Logger) hendler.Handler {
	return &handler{
		logger: logger,
	}

}
/*фуникция для реализации интерфейса handler */
func (h* handler)Register(router* httprouter.Router)  {
	router.GET("/taxi",h.GetTaxi)//добавление taxi
	router.GET("/taxi/:id",h.GetTaxiByUUID)//добавление юзера UUID
	router.POST("/taxi/:id ",h.CreateTaxi)//создание taxi
	router.PUT("/taxi/:id",h.UpdateTaxi)//обновление юзера
	router.PATCH("/taxi/id ",h.PartiallyUpdateTaxi)//частичное обновление юзера
	router.DELETE("/taxi/id ",h.DeleteTaxi)//удаление юзера
}

func (h handler) GetTaxi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (h *handler) GetTaxiByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
func (h handler) CreateTaxi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (h *handler) UpdateTaxi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
func (h handler) PartiallyUpdateTaxi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (h *handler) DeleteTaxi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}



package advert

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zhashkevych/todo-app/modules/hendler"
	_ "github.com/zhashkevych/todo-app/modules/hendler"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"net/http"
)

var _ hendler.Handler = &handler{} //подсказка реализует ли структура интерфейс с помощью Register

const (
	advertsUrl = "advert/"
)


type handler struct {
	logger loging.Logger
}

func NewHandlerAdvert(logger loging.Logger) hendler.Handler {
	return &handler{
		logger: logger,
	}

}
/*фуникция для реализации интерфейса handler */
func (h* handler)Register(router* httprouter.Router)  {

	router.GET("/users/advert",h.RegistrAdvert)//добавление обьявления
	router.POST("/users/advert/:id ",h.CreateAdvert)//создание обьявления
	router.PUT("/users/advert/:id",h.UpdateAdvert)//обновление обьявления
	router.PATCH("/users/advert/id ",h.PartiallyUpdateAdvert)//частичное обновление обьявления
	router.DELETE("/users/advert/id ",h.DeleteAdvert)//удаление обьявления
}
/*добавление юзера*/
func (h* handler)RegistrAdvert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление обьявления"))
}

/*добавление юзера UUID*/
func (h* handler)CreateAdvert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("создание обьявления а"))
}

/*создание юзера*/
func (h* handler)UpdateAdvert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("обновление обьявления"))
}

/*обновление юзера*/
func (h* handler)PartiallyUpdateAdvert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("частичное обновление обьявления"))
}

/*удаление юзера*/
func (h* handler)DeleteAdvert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("удаление обьявления"))
}
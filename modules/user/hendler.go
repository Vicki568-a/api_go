package user

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zhashkevych/todo-app/modules/hendler"
	_ "github.com/zhashkevych/todo-app/modules/hendler"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"net/http"
)

var _ hendler.Handler = &handler{} //подсказка реализует ли структура интерфейс с помощью Register

const (

	usersUrl = "users/"
	userUrl = "users/:uuid"
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
	router.GET("/users",h.GetLogin)//добавление юзера
	router.GET("/users/:id",h.GetLoginByUUID)//добавление юзера UUID
	router.POST("/users/:id ",h.CreateUser)//создание юзера
	router.PUT("/users/:id",h.UpdateUser)//обновление юзера
	router.PATCH("/users/id ",h.PartiallyUpdateUser)//частичное обновление юзера
	router.DELETE("/users/id ",h.DeleteUser)//удаление юзера
}


/*добавление юзера*/
func (h* handler)GetLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление юзера"))
}
/*добавление юзера*/
func (h* handler)GetLogin1(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление юзера"))
}
/*добавление юзера UUID*/
func (h* handler)GetLoginByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление юзер а"))
}

/*создание юзера*/
func (h* handler)CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("создание юзера"))
}

/*обновление юзера*/
func (h* handler)UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("обновление юзера"))
}

/*частичное обновление юзера*/
func (h* handler)PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("частичное обновление юзера"))
}

/*удаление юзера*/
func (h* handler)DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("удаление юзера"))
}
package future_chats_folder

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zhashkevych/todo-app/modules/hendler"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"net/http"
)

var _ hendler.Handler = &handler{} //подсказка реализует ли структура интерфейс с помощью Register

type handler struct {
	logger loging.Logger
}

func NewHandlerChat(logger loging.Logger) hendler.Handler {
	return &handler{
		logger: logger,
	}

}
/*фуникция для реализации интерфейса handler */
func (h*handler)Register(router* httprouter.Router)  {

	router.GET("/chat",h.GetLoginChat)//добавление юзера
	router.POST("/chat/:id ",h.CreateUserChat)//создание юзера
	router.PUT("/chat/:id",h.UpdateUserChat)//обновление юзера
	router.PATCH("/chat/id ",h.PartiallyUpdateUserChat)//частичное обновление юзера
	router.DELETE("/chat/id ",h.DeleteUserChat)//удаление юзера
}
/*добавление юзера*/
func (h*handler)GetLoginChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление chat"))
}

/*добавление юзера UUID*/
func (h*handler)GetLoginByUUIDChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("добавление chat а"))
}

/*создание юзера*/
func (h*handler)CreateUserChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("создание chat"))
}

/*обновление юзера*/
func (h*handler)UpdateUserChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("обновление chat"))
}

/*частичное обновление юзера*/
func (h*handler)PartiallyUpdateUserChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("частичное обновление chat"))
}

/*удаление юзера*/
func (h*handler)DeleteUserChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("удаление юзера"))
}
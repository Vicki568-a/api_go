package hendler


import (
	"github.com/julienschmidt/httprouter"
)
/*создаиние основниго хендлера*/
type Handler interface {

	Register(router* httprouter.Router)//ссылка на роутер

}


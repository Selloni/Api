package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
}

func (h *handler) Register(route *httprouter.Router) {
	route.GET("/user", GetList)
}

func (h *handler) GetList(w http.ResponseWriter, r http.Request, param httprouter.Params) {

}

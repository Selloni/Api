package user

import (
	"RestApi/Rest/interal/handlers"
	"RestApi/Rest/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//var _ handlers.Handler = &handler{}

const (
	UrlUser = "/user/:id/"
	UrlList = "/user/"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(UrlList, h.GetList)
	router.GET(UrlUser, h.GetUser)
	router.PUT(UrlUser, h.UpdateUser)
	router.DELETE(UrlUser, h.DeleteUser)
	router.POST(UrlUser, h.CrateUser)
	//router.PATCH(UrlUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("Get all list")); err != nil {
		log.Fatal(err)
	}
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Get User"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Update user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Delete user"))
}
func (h *handler) CrateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Create"))
}

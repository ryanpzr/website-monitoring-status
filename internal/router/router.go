package router

import (
	"internal/controller"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/sites", controller.PostSite).Methods("POST")

	return r
}

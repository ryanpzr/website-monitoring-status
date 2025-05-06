package router

import (
	"website-monitoring/internal/controller"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/sites", controller.PostSite).Methods("POST")
	r.HandleFunc("/sites", controller.GetAllSites).Methods("GET")
	r.HandleFunc("/sites/{id}", controller.GetSiteById).Methods("GET")
	r.HandleFunc("/history-checks/all", controller.GetAllChecksHistory).Methods("GET")
	return r
}

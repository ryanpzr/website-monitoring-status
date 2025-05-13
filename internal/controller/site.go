package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"website-monitoring/internal"
	"website-monitoring/internal/model"
	"website-monitoring/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func PostSite(w http.ResponseWriter, r *http.Request) {
	var siteInformation model.Site

	err := json.NewDecoder(r.Body).Decode(&siteInformation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(siteInformation)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errMap = make(map[string]string)
		for i := range validationErrors {
			internal.GetMessageFromFieldError(&errMap, validationErrors[i])
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMap)
		return
	}

	siteCreated, err := service.PostSite(siteInformation)
	if err != nil {
		http.Error(w, "Erro ao criar site: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao criar site: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(siteCreated)
}

func GetSiteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	site, err := service.GetSiteById(id)
	if err != nil {
		http.Error(w, "Erro ao buscar status do site: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao buscar status do site: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(site)

}

func GetAllSites(w http.ResponseWriter, r *http.Request) {
	siteList, err := service.GetAllSites()
	if err != nil {
		http.Error(w, "Erro ao buscar sites: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao buscar sites: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(siteList)
}

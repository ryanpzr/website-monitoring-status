package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"website-monitoring/internal/model"
	"website-monitoring/internal/service"
)

func PostSite(w http.ResponseWriter, r *http.Request) {
	var siteInformation model.Site

	err := json.NewDecoder(r.Body).Decode(&siteInformation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

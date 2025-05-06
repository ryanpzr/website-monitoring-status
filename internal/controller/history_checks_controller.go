package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"website-monitoring/internal/service"
)

func GetAllChecksHistory(w http.ResponseWriter, r *http.Request) {
	checksList, err := service.GetAllChecksHistory()
	if err != nil {
		http.Error(w, "Erro ao buscar sites: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao buscar sites: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checksList)
}

package controller

import (
	"log"
	"net/http"
	"website-monitoring/internal/service"

	"github.com/gin-gonic/gin"
)

func GetAllChecksHistory(c *gin.Context) {
	checksList, err := service.GetAllChecksHistory()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao buscar o histórico de checks. " + err.Error()})
		log.Printf("Erro ao buscar sites: %v", err)
		return
	}

	c.JSON(http.StatusOK, checksList)
}

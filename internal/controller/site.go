package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"website-monitoring/internal"
	"website-monitoring/internal/model"
	"website-monitoring/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func PostSite(c *gin.Context) {
	var siteInformation model.Site
	err := c.ShouldBindJSON(&siteInformation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possivel extrair a informação do body. " + err.Error()})
		return
	}

	err = validate.Struct(siteInformation)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errMap = make(map[string]string)
		for i := range validationErrors {
			internal.GetMessageFromFieldError(&errMap, validationErrors[i])
		}
		var sb strings.Builder
		for field, msg := range errMap {
			sb.WriteString(fmt.Sprintf("%s: %s\n", field, msg))
		}
		c.JSON(http.StatusBadRequest, sb.String())
		return
	}

	siteCreated, err := service.PostSite(siteInformation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar o site. " + err.Error()})
		log.Printf("Erro ao criar site: %v", err)
		return
	}

	c.JSON(http.StatusAccepted, siteCreated)
}

func GetSiteById(c *gin.Context) {
	// Vars pega apenas a informação passada na url da requisição.
	id := c.Param("id")

	site, err := service.GetSiteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro buscar site por ID. " + err.Error()})
		log.Printf("Erro ao buscar status do site: %v", err)
		return
	}

	c.JSON(http.StatusOK, site)
}

func GetAllSites(c *gin.Context) {
	siteList, err := service.GetAllSites()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro buscar todos os sites. " + err.Error()})
		log.Printf("Erro ao buscar sites: %v", err)
		return
	}

	c.JSON(http.StatusOK, siteList)
}

package router

import (
	"website-monitoring/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sites", controller.PostSite)
	r.GET("/sites", controller.GetAllSites)
	r.GET("/sites/:id", controller.GetSiteById)
	r.GET("/history-checks/all", controller.GetAllChecksHistory)
	return r
}

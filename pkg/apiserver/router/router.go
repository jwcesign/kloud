package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jwcesign/kloud/pkg/apiserver/handler"
)

func NewAnalyzerAPIServer() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	corsHandler := cors.New(config)
	router.Use(corsHandler)

	initRouter(router)

	return router
}

func initRouter(router *gin.Engine) {
	group := router.Group("api/v1")
	group.GET("/migration/configuration", nil)
	group.POST("/migration/configuration", nil)
	group.GET("/migration/costs", nil)
	group.POST("/migration/costs", nil)

	group.GET("/healthz", handler.HealthzHandler)
}

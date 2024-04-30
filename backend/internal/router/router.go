package router

import (
	"sn-go-api/internal/api"
	"sn-go-api/internal/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(snConfig *config.ServiceNowConfig) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/incidents", func(c *gin.Context) {
				api.FetchIncidents(c, snConfig)
		})
	}
	return router
}

package router

import (
	"sn-go-api/internal/api"
	"sn-go-api/internal/config"

	"github.com/gin-gonic/gin"
)

func getIncidents(snConfig *config.ServiceNowConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.FetchIncidents(c, snConfig)
	}
}

func SetupRouter(snConfig *config.ServiceNowConfig) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/incidents", getIncidents(snConfig))
	}
	return router
}

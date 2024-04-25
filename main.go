package main

import (
	"sn-go-api/internal/api"
	"sn-go-api/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	snConfig := config.Init()
	r := setupRouter(snConfig)
	r.Run(":8080")
}

func setupRouter(snConfig *config.ServiceNowConfig) *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/incidents", func(c *gin.Context) {
				api.FetchIncidents(c, snConfig)
		})
	}
	return r
}

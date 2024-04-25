package main

import (
	"encoding/json"
	"net/http"

	"sn-go-api/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	snConfig := config.Init()
	r := setupRouter(snConfig)
	r.Run(":8080")
}

func setupRouter(snConfig *config.ServiceNowConfig) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/incidents", func(c *gin.Context) {
				fetchFromServiceNow(c, snConfig)
			})
		}
	}
	return r
}

func buildURL(instance string) string {
	return "https://" + instance + ".service-now.com/api/now/table/incident"
}

func fetchFromServiceNow(c *gin.Context, snConfig *config.ServiceNowConfig) {
	client := resty.New()
	client.SetBasicAuth(snConfig.Username, snConfig.Password)
	userQuery := c.DefaultQuery("filter", "active=true")
	limit := c.DefaultQuery("limit", "10")
	fields := c.DefaultQuery("fields", "active,assigned_to,number,short_description,priority,sys_id")

	fullURL := buildURL(snConfig.Instance)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParam("sysparm_query", userQuery).
		SetQueryParam("sysparm_limit", limit).
		SetQueryParam("sysparm_fields", fields).
		Get(fullURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from ServiceNow"})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

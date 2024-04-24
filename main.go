package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/incidents", fetchFromServiceNow)
	return r
}

func fetchFromServiceNow(c *gin.Context) {
	client := resty.New()
	client.SetBasicAuth("user", "******")
	userQuery := c.DefaultQuery("filter", "active=true")
	limit := c.DefaultQuery("limit", "10")
	fields := c.DefaultQuery("fields", "active,assigned_to,number,short_description,priority,sys_id")

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParam("sysparm_query", userQuery).
		SetQueryParam("sysparm_limit", limit).
		SetQueryParam("sysparm_fields", fields).
		Get("https://instance.service-now.com/api/now/table/incident")

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

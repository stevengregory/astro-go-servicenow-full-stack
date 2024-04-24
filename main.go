package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	r := setupRouter()
	r.Run(":8080")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/incidents", fetchFromServiceNow)
	return r
}

func fetchFromServiceNow(c *gin.Context) {
	client := resty.New()
	username := viper.GetString("servicenow.username")
	password := viper.GetString("servicenow.password")
	baseURL := viper.GetString("servicenow.baseURL")

	client.SetBasicAuth(username, password)
	userQuery := c.DefaultQuery("filter", "active=true")
	limit := c.DefaultQuery("limit", "10")
	fields := c.DefaultQuery("fields", "active,assigned_to,number,short_description,priority,sys_id")

	fullURL := baseURL + "/api/now/table/incident"
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

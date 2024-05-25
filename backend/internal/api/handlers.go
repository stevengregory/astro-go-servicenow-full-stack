package api

import (
	"encoding/json"
	"net/http"

	"sn-go-api/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func buildURL(instance string) string {
	return "https://" + instance + ".service-now.com/api/now/table/incident"
}

func configureClient(username, password string) *resty.Client {
	client := resty.New()
	client.SetBasicAuth(username, password)
	return client
}

func doRequest(client *resty.Client, url, userQuery, limit, fields string) (*resty.Response, error) {
	return client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParam("sysparm_query", userQuery).
		SetQueryParam("sysparm_limit", limit).
		SetQueryParam("sysparm_fields", fields).
		SetQueryParam("sysparm_display_value", "true").
		Get(url)
}

func parseResponse(resp *resty.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func prepQueryParams(c *gin.Context) (string, string, string) {
	userQuery := c.DefaultQuery("filter", "active=true")
	limit := c.DefaultQuery("limit", "10")
	fields := c.DefaultQuery("fields", "active,assigned_to,number,short_description,priority,sys_id")
	return userQuery, limit, fields
}

func FetchIncidents(c *gin.Context, snConfig *config.ServiceNowConfig) {
	client := configureClient(snConfig.Username, snConfig.Password)
	userQuery, limit, fields := prepQueryParams(c)
	url := buildURL(snConfig.Instance)
	resp, err := doRequest(client, url, userQuery, limit, fields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from ServiceNow"})
		return
	}
	result, err := parseResponse(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}
	c.JSON(http.StatusOK, result)
}

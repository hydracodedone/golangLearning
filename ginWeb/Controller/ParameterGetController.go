package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStringParameterHandler(c *gin.Context) {
	parameter, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"message": "Parameter dose not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Parameter is %v", parameter)})
}

func FormParameterHandler(c *gin.Context) {
	parameter, ok := c.GetPostForm("name")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"message": "Parameter dose not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Parameter is %v", parameter)})
}

func JsonParameterHandler(c *gin.Context) {
	parameter, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Parameter read fail"})
		return
	}
	var mapParamter map[string]interface{}
	err = json.Unmarshal(parameter, &mapParamter)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Parameter unmarshal fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Parameter is %v", mapParamter)})
}

package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogHandler(c *gin.Context) {
	logrus.Info(c.Request.Method, c.Request.URL)
	c.JSON(http.StatusOK, gin.H{"message": "logging successfully"})
}

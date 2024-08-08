package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminLogin godoc
// @Summary SwaggerDemo
// @Description SwaggerDemo
// @Tags Swagger
// @ID /SwaggerDemo/Demo1
// @Accept  json
// @Produce  json
// @Router /SwaggerDemo/Demo1 [get]
func SwaggerDemoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

package Controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AsyncHandler(c *gin.Context) {
	copyCtx := c.Copy()
	go AsyncFunction(copyCtx)
	fmt.Println("Asynchronous task begin")
	c.JSON(http.StatusOK, gin.H{"message": "Asynchronous task begin"})
}
func AsyncFunction(c *gin.Context) {
	time.Sleep(3 * time.Second)
	fmt.Printf("the context is %#v\n", c.Request.Host)
}

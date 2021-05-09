package Controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookieHandler(c *gin.Context) {
	_, err := c.Cookie("extra-info")
	if err != nil {
		c.SetCookie("extra-info", "Hydra", 60, "/Cookie", "*", false, false)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Set Cookie successfully"})
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("extra-info")
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/Cookie/SetCookie")
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Get Cookie :[%v] successfully", cookie)})
}

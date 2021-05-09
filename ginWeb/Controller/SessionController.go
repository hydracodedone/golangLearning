package Controller

import (
	"fmt"
	"net/http"

	"demo_for_gin/Utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("count", 0)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Set Session successfully"})
}

func GetSessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("count")
	if value == nil {
		c.Redirect(http.StatusMovedPermanently, "/Session/SetSession")
		return
	} else {
		count := value.(int)
		count++
		fmt.Printf("now count is %d\n", count)
		session.Set("count", count)
		session.Save()
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Session successfully"})
}

func SetSessionHandler2(c *gin.Context) {
	session, err := Utils.Store.Get(c.Request, "login_info")
	if err != nil {
		fmt.Println("generate session fails")
		c.JSON(http.StatusOK, gin.H{"message": "Set Session2 fail"})
		return
	}
	_, ok := session.Values["count"]
	if ok {
		c.JSON(http.StatusOK, gin.H{"message": "Do not need set Session2"})
		return
	}
	session.Values["count"] = 0
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Set Session2 successfully"})
}

func GetSessionHandler2(c *gin.Context) {
	session, err := Utils.Store.Get(c.Request, "login_info")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/Session/SetSession2")
		return
	} else {
		count := session.Values["count"]
		fmt.Printf("now count is %v\n", count)
		session.Values["count"] = count.(int) + 1
		session.Save(c.Request, c.Writer)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Session2 successfully"})
}

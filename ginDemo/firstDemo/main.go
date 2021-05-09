package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func demo() {
	engine := gin.Default()
	engine.GET("/v1/HelloWorld", helloWorld)
	engine.GET("/v1/health", healthCheck)

	err := engine.Run(":9000")
	if err != nil {
		log.Fatalf("Gin Engine Run Fail:<%s>", err.Error())
	}
}
func healthCheck(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": 200,
	})
}
func helloWorld(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "HELLO,WORLD",
	})
}

func main() {
	demo()
}

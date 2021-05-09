package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func tlsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "443",
		})
		err := secureMiddleware.Process(context.Writer, context.Request)
		if err != nil {
			log.Fatalf("The secureMiddleware handle fail:<%s>", err.Error())
		}
		context.Next()
	}
}

func demo() {
	engine := gin.Default()
	engine.Use(tlsHandler())
	engine.GET("/home", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	err := engine.RunTLS(":8000", "/home/hydra/Project/GolangLearning/src/ginDemo/httpAndHttpsDemo/certification", "/home/hydra/Project/GolangLearning/src/ginDemo/httpAndHttpsDemo/certification/")
	if err != nil {
		log.Fatalf("Run TLS Fail:<%s>", err.Error())
	}
}

func main() {
	demo()
}

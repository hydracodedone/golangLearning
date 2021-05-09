package main

import (
	"fmt"

	"demo_for_gin/Middleware"
	"demo_for_gin/Router"
	"demo_for_gin/Utils"
	"demo_for_gin/Validator"

	"github.com/gin-gonic/gin"
)

func init() {
	Utils.LoggerHookTest()

}
func main() {
	r := gin.New()
	Middleware.LoadAll(r)
	Router.LoadAll(r)
	Validator.LoadAll(r)
	if err := r.Run("0.0.0.0:8888"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}

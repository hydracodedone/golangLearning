package Utils

import "github.com/gin-gonic/gin"

func RegisterMiddleware(handler gin.HandlerFunc, middlewreHandler ...gin.HandlerFunc) []gin.HandlerFunc {
	if len(middlewreHandler) == 0 {
		var handlerSlice []gin.HandlerFunc = make([]gin.HandlerFunc, 1)
		handlerSlice[0] = handler
		return handlerSlice
	}
	middlewreHandler = append(middlewreHandler, handler)
	return middlewreHandler
}

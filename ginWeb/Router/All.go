package Router

import "github.com/gin-gonic/gin"

func LoadAll(e *gin.Engine) {
	loadBinding(e)
	loadAsync(e)
	loadCookie(e)
	loadSession(e)
	loadLog(e)
	loadParamterGet(e)
	loadSwagger(e)
}

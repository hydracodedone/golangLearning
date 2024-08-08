package Middleware

import (
	limit "github.com/aviddiviner/gin-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	stats "github.com/semihalev/gin-stats"
)

func LoadAll(e *gin.Engine) {
	loadSessionMiddleware(e)
	e.Use(gin.Recovery())
	e.Use(GlobalMiddleware())
	e.Use(limit.MaxAllowed(20))
	e.Use(helmet.Default())
	e.Use(stats.RequestStats())
}

func loadSessionMiddleware(e *gin.Engine) {
	sessionStoreEngine := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("session", sessionStoreEngine))
}

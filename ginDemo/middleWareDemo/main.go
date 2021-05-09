package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
https://github.com/gin-gonic/contrib/blob/master/README.md
RestGate - Secure authentication for REST API endpoints
staticbin - middleware/handler for serving static files from binary data
gin-cors - Official CORS gin's middleware
gin-csrf - CSRF protection
gin-health - middleware that simplifies stat reporting via gocraft/health
gin-merry - middleware for pretty-printing merry errors with context
gin-revision - Revision middleware for Gin framework
gin-jwt - JWT Middleware for Gin Framework
gin-sessions - session middleware based on mongodb and mysql
gin-location - middleware for exposing the server's hostname and scheme
gin-nice-recovery - panic recovery middleware that lets you build a nicer user experience
gin-limiter - A simple gin middleware for ip limiter based on redis.
gin-limit - limits simultaneous requests; can help with high traffic load
gin-limit-by-key - An in-memory middleware to limit access rate by custom key and rate.
ez-gin-template - easy template wrap for gin
gin-hydra - Hydra middleware for Gin
gin-glog - meant as drop-in replacement for Gin's default logger
gin-gomonitor - for exposing metrics with Go-Monitor
gin-oauth2 - for working with OAuth2
static An alternative static assets handler for the gin framework.
xss-mw - XssMw is a middleware designed to "auto remove XSS" from user submitted input
gin-helmet - Collection of simple security middleware.
gin-jwt-session - middleware to provide JWT/Session/Flashes, easy to use while also provide options for adjust if necessary. Provide sample too.
gin-template - Easy and simple to use html/template for gin framework.
pongo2gin - Package pongo2gin is a template renderer that can be used with the Gin web framework [pongo2 like django templates]
gin-redis-ip-limiter - Request limiter based on ip address. It works with redis and with a sliding-window mechanism.
gin-method-override - Method override by POST form param _method, inspired by Ruby's same name rack
gin-access-limit - An access-control middleware by specifying allowed source CIDR notations.
gin-session - Session middleware for Gin
gin-stats - Lightweight and useful request metrics middleware
gin-statsd - A Gin middleware for reporting to statsd deamon
gin-health-check - A health check middleware for Gin
gin-session-middleware - A efficient, safely and easy-to-use session library for Go.
ginception - Nice looking exception page
gin-inspector - Gin middleware for investigating http request.
gin-dump - Gin middleware/handler to dump header/body of request and response. Very helpful for debugging your applications.
go-gin-prometheus - Gin Prometheus metrics exporter
ginprom - Prometheus metrics exporter for Gin
gin-go-metrics - Gin middleware to gather and store metrics using rcrowley/go-metrics
ginrpc - Gin middleware/handler auto binding tools. support object register by annotated route like beego
goscope - Watch incoming requests, outgoing responses and logs of your Gin application with this plug and play middleware inspired by Laravel Telescope.
gin-nocache - NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent a router (or subrouter) from being cached by an upstream proxy and/or client.
logging - logging provide GinLogger uses zap to log detailed access logs in JSON or text format with trace id, supports flexible and rich configuration, and supports automatic reporting of log events above error level to sentry
ratelimiter - Gin middleware for token bucket ratelimiter.
servefiles - serving static files with performance-enhancing cache control headers; also handles gzip & brotli compressed files
*/
func demo() {
	engine := gin.Default()
	globalMiddleware := func(context *gin.Context) {
		log.Println("globalMiddleware in")
		context.Next()
		log.Println("globalMiddleware out")
	}
	globalMiddleware2 := func(context *gin.Context) {
		log.Println("globalMiddleware2 in")
		context.Next()
		log.Println("globalMiddleware2 out")
	}
	engine.Use(globalMiddleware, globalMiddleware2)
	localMiddleware := func(context *gin.Context) {
		log.Println("localMiddleware in")
		log.Println("localMiddleware out")
	}
	engine.GET("/home", localMiddleware, func(context *gin.Context) {
		log.Println("home in")
		context.JSON(200, gin.H{
			"message": 200,
		})
		log.Println("home out")

	})
	err := engine.Run(":9000")
	if err != nil {
		log.Fatalf("Gin Engine Run Fail:<%s>", err.Error())
	}
}
func main() {
	demo()
}

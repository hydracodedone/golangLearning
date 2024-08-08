package main

import (
	"distribute_tracing_gin/middleware"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func getOrInitSpanByContext(c *gin.Context, funcName string) opentracing.Span {
	tracer, exsist := c.Get("tracer")
	if !exsist {
		panic("no tracer found")
	}
	realTracer := tracer.(opentracing.Tracer)
	parentSpanCtx, exsist := c.Get("spanCtx")
	var currentSpan opentracing.Span
	if exsist {
		realParentSpanCtx := parentSpanCtx.(opentracing.SpanContext)
		currentSpan = realTracer.StartSpan(funcName, opentracing.ChildOf(realParentSpanCtx))
	} else {
		currentSpan = realTracer.StartSpan(funcName)
	}
	c.Set("spanCtx", currentSpan.Context())
	return currentSpan
}

type Response struct {
	Message string `json:"message"`
}

func crossRequestTwice(c *gin.Context) {
	fmt.Println(c.Request.Header)
	currentSpan := getOrInitSpanByContext(c, "crossRequestTwice")
	currentSpan.LogFields(
		log.String("currentFunc", "crossRequestTwice"),
	)
	defer currentSpan.Finish()
	resp := &Response{Message: "crossRequestTwice"}
	c.JSON(200, resp)
}
func crossRequest(c *gin.Context) {
	currentSpan := getOrInitSpanByContext(c, "crossRequest")
	currentSpan.LogFields(
		log.String("currentFunc", "crossRequest"),
	)
	defer currentSpan.Finish()
	resp := &Response{Message: "crossRequest"}
	func() {
		req, err := http.NewRequest("GET", "http://localhost:8000/crossRequestTwice", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = currentSpan.Tracer().Inject(currentSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		if err != nil {
			panic(err)
		}
		client := &http.Client{}
		reqResp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer reqResp.Body.Close()
		_, err = io.ReadAll(reqResp.Body)
		if err != nil {
			panic(err)
		}
	}()
	c.JSON(200, resp)
}

func once(c *gin.Context) {
	currentSpan := getOrInitSpanByContext(c, "once")
	currentSpan.LogFields(
		log.String("currentFunc", "once"),
	)
	defer currentSpan.Finish()
	resp := &Response{Message: "onceFunc"}
	c.JSON(200, resp)
}
func twiceCal(c *gin.Context) int {
	currentSpan := getOrInitSpanByContext(c, "twiceCal")
	currentSpan.LogFields(
		log.String("currentFunc", "twiceCalFunc"),
	)
	defer currentSpan.Finish()
	return 100
}
func twice(c *gin.Context) {
	currentSpan := getOrInitSpanByContext(c, "twice")
	currentSpan.LogFields(
		log.String("currentFunc", "twiceFunc"),
	)
	defer currentSpan.Finish()
	twiceCal(c)
	resp := &Response{Message: "twice"}
	c.JSON(200, resp)
}
func main() {
	r := gin.Default()
	r.Use(middleware.TraceMiddleWare())
	r.GET("/once", once)
	r.GET("/twice", twice)
	r.GET("/crossRequest", crossRequest)
	r.GET("/crossRequestTwice", crossRequestTwice)

	r.Run("localhost:8000")
}

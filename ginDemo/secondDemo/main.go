package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func customUserNameValidate(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return value != ""
	}
	return false
}

func returnEngine() *gin.Engine {
	jwtSecretKey := []byte(time.Now().String())
	type JwtStruct struct {
		UserId uint
		jwt.StandardClaims
	}
	engine := gin.New()
	file, _ := os.Create("./ginLog.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // write log both to stdout and log file
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	version1Group := engine.Group("/v1", func(context *gin.Context) {
		log.Printf("Enter v1 route group")
	})
	version1Group.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": 200,
		})
	})
	version1Group.GET("/getMethod", func(context *gin.Context) {
		query := context.Query("info")
		defaultQuery := context.DefaultQuery("info", "default")
		getQuery, ok := context.GetQuery("info")
		log.Printf("The query is %v\n", query)
		log.Printf("The defaultQuery is %v\n", defaultQuery)
		log.Printf("The getQuery is %v,The ok is %v\n", getQuery, ok)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	version1Group.POST("/postMethod", func(context *gin.Context) {
		postForm := context.PostForm("info")
		defaultPostForm := context.DefaultPostForm("info", "default")
		getPostForm, ok := context.GetPostForm("info")
		log.Printf("The postForm is %v\n", postForm)
		log.Printf("The defaultPostForm is %v\n", defaultPostForm)
		log.Printf("The getPostForm is %v,The ok is %v\n", getPostForm, ok)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	/*
		if we use /:name/:age we only match /hydra/23,but we can not match /hydra /hydra/ /hydra/23/
		if we use /:name/*age we can  match /hydra/ /hydra/23/2 /hydra/23 /hydra/23/
	*/
	version1Group.GET("/urlParameter/:name/*age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		log.Printf("The name is %v\n", name)
		log.Printf("The age is %v\n", age)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	version1Group.POST("/formBind", func(context *gin.Context) {
		type UserInfo struct {
			UserName string `form:"name" binding:"required"`
			UserAge  int    `form:"age" binding:"required"`
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		userInstance := new(UserInfo)
		err := context.ShouldBind(userInstance)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			log.Printf("The User Info is %v\n", userInstance)
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "ok",
			})
		}
	})
	version1Group.POST("/jsonBind", func(context *gin.Context) {
		type UserInfo struct {
			UserName string `json:"name" binding:"required"`
			UserAge  int    `json:"age" binding:"required"`
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		userInstance := new(UserInfo)
		err := context.ShouldBindJSON(userInstance)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			log.Printf("The User Info is %v\n", userInstance)
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "ok",
			})
		}
	})
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("customUserNameValidate", customUserNameValidate)
		if err != nil {
			log.Fatalf("Bind custom validation fail:<%s>", err.Error())
		} else {
			log.Println("Bind custom validation success")
		}
	} else {
		log.Fatalf("custom validation initial fail")
	}
	version1Group.POST("/customBind", func(context *gin.Context) {
		type UserInfo struct {
			UserName string `json:"name" binding:"customUserNameValidate"`
			UserAge  int    `json:"age" binding:"required"`
		}
		userInstance := new(UserInfo)
		err := context.ShouldBindJSON(userInstance)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			log.Printf("The User Info is %v\n", userInstance)
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "ok",
			})
			return
		}
	})
	version1Group.GET("/temporaryRedirect", func(context *gin.Context) {
		url := context.Request.URL
		path := url.Path
		redirectUrl := "/" + strings.Split(path, "/")[1] + "/" + "redirectDestination"
		log.Printf("The PATH is %s\n", path)
		log.Printf("The Redirect Path is %s\n", redirectUrl)
		context.Request.URL.Path = redirectUrl
		engine.HandleContext(context)
	})
	version1Group.GET("/permanentRedirect", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "https://www.baidu.com")
	})
	version1Group.GET("/redirectDestination", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	version1Group.GET("/setCookie", func(context *gin.Context) {
		context.SetCookie(
			"name",
			"hydra",
			3600,
			"/v1",
			"localhost",
			false,
			true,
		)
		log.Printf("Set cookie success")
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	version1Group.GET("/getCookie", func(context *gin.Context) {
		cookie, err := context.Cookie("name")
		if err != nil {
			log.Printf("Get cookie fail:<%s>", err.Error())
			context.Request.URL.Path = "/v1/setCookie"
			engine.HandleContext(context)
		} else {
			log.Printf("The cookie is %s\n", cookie)
		}
	})
	version1Group.GET("/setJwtToken", func(context *gin.Context) {
		userId, ok := context.GetQuery("userId")
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "request must contain user id",
			})
			return
		}
		userIdToInt, _ := strconv.Atoi(userId)
		userIdToUint := uint(userIdToInt)
		jwtValue := &JwtStruct{
			UserId: userIdToUint,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "localhost",
				Subject:   "user token",
			},
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtValue)
		tokenString, err := jwtToken.SignedString(jwtSecretKey)
		if err != nil {
			log.Fatalf("Jwt token generate fail:<%s>", err.Error())
			return
		} else {
			log.Printf("Jwt token generage success: %s\n", tokenString)
			context.Header("Authentication", tokenString)
			context.SetCookie(
				"Authentication",
				tokenString, 3600,
				"/v1",
				"localhost",
				false,
				true,
			)
			context.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			return
		}
	})

	version1Group.GET("/getJwtToken", func(context *gin.Context) {
		tokenString, err := context.Cookie("Authentication")
		if err != nil {
			log.Printf("Get cookie fail:<%s>", err.Error())
			context.Request.URL.Path = "/v1/setJwtToken"
			engine.HandleContext(context)
			return
		}
		if tokenString == "" {
			context.Request.URL.Path = "/v1/setJwtToken"
			engine.HandleContext(context)
			return
		} else {
			jwtValue := &JwtStruct{}
			jwtToken, ok := jwt.ParseWithClaims(tokenString, jwtValue, func(token *jwt.Token) (interface{}, error) {
				return jwtSecretKey, nil
			})
			if ok != nil {
				log.Fatalln("Jwt token parse fail")
			} else {
				if !jwtToken.Valid {
					context.JSON(http.StatusBadRequest, gin.H{
						"message": "JwtToken is not Valid",
					})
					return
				} else {
					context.JSON(http.StatusOK, gin.H{
						"message": "ok",
						"userId":  jwtValue.UserId,
					})
					return
				}
			}
		}
	})

	version1Group.GET("/async", func(context *gin.Context) {
		log.Printf("Async Task Begin")
		copyContext := context.Copy()
		go func() {
			time.Sleep(time.Second)
			log.Printf("The Request Path is %s\n", copyContext.Request.URL.Path)
			log.Printf("Async Task is Done")
		}()
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	version1Group.Any("/anyMethod", func(context *gin.Context) {
		log.Printf("The Request Method is %s\n", context.Request.Method)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "404",
		})
	})
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "405",
		})
	})
	return engine
}

func demo() {
	engine := returnEngine()
	err := http.ListenAndServe(":9000", engine)
	if err != nil {
		log.Fatalf("Http Listen And Server Fail:<%s>", err.Error())
	}
}
func demo2() {
	engine := returnEngine()
	s := &http.Server{
		Addr:           ":9000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Http Listen And Server Fail:<%s>", err.Error())
	}
}
func main() {
	demo2()
}

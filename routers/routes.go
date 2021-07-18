package routers

import (
	"net/http"
	
	"boframe/controller/user"
	"boframe/logger"
	"boframe/middleware"
	"github.com/gin-gonic/gin"
)

func checkGinMode(mode string) (res string) {
	switch mode {
	case "debug":
		res = gin.DebugMode
	case "test":
		res = gin.TestMode
	case "release", "production":
		res = gin.ReleaseMode
	default:
		res = gin.DebugMode
	}
	return
}

func Setup(mode string) *gin.Engine {
	gin.SetMode(checkGinMode(mode))
	r := gin.New()
	
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
	}
	
	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})
	}
	
	return r
}

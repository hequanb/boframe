package middleware

import (
	"net/http"
	"time"
	
	"boframe/controller"
	"boframe/pkg/errcode"
	"boframe/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ContextUserIdKey   = "user_id"
	ContextUsernameKey = "username"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			controller.ResponseErrorCodeWithHTTPCode(ctx, http.StatusUnauthorized, errcode.Build(errcode.ErrNotLogin))
			ctx.Abort()
			return
		}
		
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			zap.L().Error("auth failed, jwt parse token failed", zap.Error(err))
			controller.ResponseErrorCodeWithHTTPCode(ctx, http.StatusUnauthorized, errcode.Build(errcode.ErrInvalidToken))
			ctx.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			controller.ResponseErrorCodeWithHTTPCode(ctx, http.StatusUnauthorized, errcode.Build(errcode.ErrTokenExpired))
			ctx.Abort()
			return
		}
		
		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		ctx.Set(ContextUserIdKey, claims.UserId)
		ctx.Set(ContextUsernameKey, claims.Username)
		
		ctx.Next()
	}
}

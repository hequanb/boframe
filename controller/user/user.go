package user

import (
	"boframe/controller"
	"boframe/pkg/errcode"
	"boframe/pkg/validator"
	"boframe/service/userservice"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type registerRequest struct {
	Username   string `json:"username" binding:"required,gte=5,lte=50"`
	Nickname   string `json:"nickname" binding:"required,gte=2,lte=50"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,gte=6,lte=30"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type registerResponse struct {
	Uid int64 `json:"uid"`
}

func Register(ctx *gin.Context) {
	requestParam := new(registerRequest)
	// 接收参数，参数校验
	if err := ctx.ShouldBindJSON(requestParam); err != nil {
		zap.L().Error("register bind failed", zap.Error(err))
		ok := validator.IsValidationErrors(err)
		if !ok {
			controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrInvalidParam))
			return
		}
		
		controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrInvalidParam))
		return
	}
	
	bo := userservice.RegisterBo{
		Username:   requestParam.Username,
		Nickname:   requestParam.Nickname,
		Email:      requestParam.Email,
		Password:   requestParam.Password,
		RePassword: requestParam.RePassword,
	}
	// 调用logic
	uid, err := userservice.Register(&bo)
	if err != nil {
		if code, ok := errcode.IsErrorCode(err); ok {
			controller.ResponseErrorCode(ctx, code)
		} else {
			controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrServer))
		}
		zap.L().Error("user service register failed", zap.Error(err))
		return
	}
	
	res := new(registerResponse)
	res.Uid = uid
	// 返回结果
	controller.ResponseSuccessWithData(ctx, res)
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	// 参数解析
	loginRequest := new(loginRequest)
	if err := ctx.ShouldBindJSON(loginRequest); err != nil {
		zap.L().Error("login bind failed", zap.Error(err))
		if ok := validator.IsValidationErrors(err); !ok {
			controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrInvalidParam))
			return
		}
		
		controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrInvalidParam))
		return
	}
	
	bo := userservice.LoginBo{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}
	token, err := userservice.Login(&bo)
	if code, ok := errcode.IsErrorCode(err); ok {
		controller.ResponseErrorCode(ctx, code)
		return
	}
	if err != nil {
		zap.L().Error("login failed",
			zap.String("username", loginRequest.Username),
			zap.Error(err),
		)
		controller.ResponseErrorCode(ctx, errcode.Build(errcode.ErrServer))
		return
	}
	
	resp := new(loginResponse)
	resp.Token = token
	controller.ResponseSuccessWithData(ctx, resp)
}

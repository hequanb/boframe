package controller

import (
	"net/http"
	
	"boframe/pkg/errcode"
	"github.com/gin-gonic/gin"
)

const (
	SuccessCode    = 1
	SuccessMessage = "成功"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func responseJSON(c *gin.Context, httpCode int, payload *ResponseData, abort bool) {
	c.JSON(httpCode, payload)
	if abort {
		c.Abort()
	}
}

func ResponseErrorCode(c *gin.Context, code *errcode.ErrCode) {
	res := &ResponseData{
		Code:    code.Code,
		Message: code.Message,
		Data:    code.Data,
	}
	responseJSON(c, http.StatusOK, res, true)
}

func ResponseErrorCodeWithHTTPCode(c *gin.Context, httpCode int, code *errcode.ErrCode) {
	res := &ResponseData{
		Code:    code.Code,
		Message: code.Message,
		Data:    code.Data,
	}
	responseJSON(c, httpCode, res, true)
}

func ResponseSuccess(c *gin.Context) {
	responseJSON(c, http.StatusOK, nil, false)
}

func ResponseSuccessWithData(c *gin.Context, payload interface{}) {
	res := &ResponseData{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    payload,
	}
	responseJSON(c, http.StatusOK, res, false)
}

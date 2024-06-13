package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
	Data any `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code int) {

	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  Code2Msg(code),
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  Code2Msg(CodeSuccess),
		Data: data,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code int, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

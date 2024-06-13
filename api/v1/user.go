package v1

import (
	"bluebell/api"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var userLg *logic.UserLg

// SignUpHandler 注册
// @Summary 注册
// @Tags 用户
// @Produce json
// @Param data body models.UserSignUpParams true "SignUpParams"
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数
	data := new(models.UserSignUpParams)
	if err := c.ShouldBindJSON(data); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			api.ResponseError(c, api.CodeInvalidParam)
		}
		api.ResponseErrorWithMsg(c, api.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 逻辑操作
	if code := userLg.SignUp(data); code != api.CodeSuccess {
		if code == api.CodeUserExist {
			api.ResponseErrorWithMsg(c, api.CodeUserExist, api.Code2Msg(code))
			return
		}
		api.ResponseError(c, api.CodeServerBusy)
		return
	}
	// 3. 返回响应
	//fmt.Println(data)
	api.ResponseSuccess(c, nil)
}

// LoginHandler 登录
// @Summary 登录
// @Tags 用户
// @Produce json
// @Param data body models.UserLoginParams true "LoginParams"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1. 获取参数
	data := new(models.UserLoginParams)
	if err := c.ShouldBindJSON(data); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			api.ResponseError(c, api.CodeInvalidParam)
		}
		api.ResponseErrorWithMsg(c, api.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//fmt.Println(data)
	// 2.业务逻辑处理
	code, token := userLg.Login(data)
	if code != api.CodeSuccess {
		zap.L().Error("Login with invalid param", zap.Error(errors.New(api.Code2Msg(code))))
		api.ResponseErrorWithMsg(c, code, "用户名或密码错误")
		return
	}
	// 3.返回响应
	api.ResponseSuccess(c, token)
}

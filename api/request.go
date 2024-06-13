package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const ContextClaimsKey = "userClaims"

// GetUserID 获取当前登录用户ID
func GetUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextClaimsKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

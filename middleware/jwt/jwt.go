package jwt

import (
	"bluebell/api"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

var mysecret = []byte("bluebell")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func CreateToken(userID int64, userName string) (string, error) {
	claims := MyClaims{
		UserID:   userID,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "bluebell",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mysecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mysecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从HTTP请求的Authorization头部获取Token
		// 假设Token格式为Bearer xxxxxxxxx...
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			api.ResponseError(c, api.CodeNeedLogin)
			c.Abort()
			return
		}
		// 解析Token字符串，分割Bearer和实际的Token
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			api.ResponseError(c, api.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		tokenString = parts[1]
		claims, err := ParseToken(tokenString)
		if err != nil {
			// 如果Token无效，返回错误并终止请求
			api.ResponseError(c, api.CodeInvalidToken)
			c.Abort()
			return
		}
		// 如果Token有效，将claims添加到Context中，供后续处理使用
		c.Set(api.ContextClaimsKey, claims.UserID)
		c.Next() // 后续可以使用c.Get(ContextClaimsKey)获取claims
	}
}

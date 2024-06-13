package router

import (
	v1 "bluebell/api/v1"
	"bluebell/logger"
	"bluebell/middleware/gin_ratelimit"
	"bluebell/middleware/jwt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"

	_ "bluebell/docs"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成生产模式
	}
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(logger.GinLogger(), logger.GinRecovery(true), gin_ratelimit.RatelimitMiddleware(2*time.Second, 1))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 注册
	r.POST("/api/v1/signup", v1.SignUpHandler)
	// 登录
	r.POST("/api/v1/login", v1.LoginHandler)

	r.GET("/api/v1/posts2", v1.GetPostListHandler2)

	v01 := r.Group("/api/v1")
	v01.Use(jwt.JWTMiddleware())
	{
		// 社区相关接口
		v01.GET("/community", v1.GetCommunityHandler)           // 获取社区列表
		v01.GET("/community/:id", v1.GetCommunityDetailHandler) // 获取社区详情
		v01.POST("/community", v1.CaretCommunityHandler)        // 创建社区
		v01.DELETE("/community/:id", v1.DeleteCommunityHandler) //

		// 文章相关接口
		v01.POST("/post", v1.CaretPostHandler) // 创建文章
		v01.GET("/post/:id", v1.GetPostDetailHandler)
		v01.GET("/posts", v1.GetPostListHandler)

		// 投票相关接口
		v01.POST("/vote", v1.PostVoteHandler)
	}

	return r
}

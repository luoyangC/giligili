package server

import (
	"giligili/api"
	"giligili/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 获取视频列表
		v1.GET("videos", api.ListVideo)

		// 获取指定视频
		v1.GET("video/:id", api.ShowVideo)
	}

	// 需要登录保护的
	v1.Use(middleware.JWTAuth())
	{
		// User Routing
		v1.GET("user/me", api.UserMe)

		// 创建视频
		v1.POST("videos", api.CreateVideo)

		// 更新指定视频
		v1.PUT("video/:id", api.UpdateVideo)

		// 删除指定接口
		v1.DELETE("video/:id", api.DeleteVideo)
	}
	return r
}

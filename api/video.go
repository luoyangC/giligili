package api

import (
	"giligili/service"
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideo 添加视频
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ListVideo(c *gin.Context) {
	c.JSON(200, "视频列表")
}

func ShowVideo(c *gin.Context)  {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg: "获取指定视频",
	})
}

func UpdateVideo(c *gin.Context)  {
	c.JSON(200, "更新视频")
}

func DeleteVideo(c *gin.Context)  {
	c.JSON(200, "删除视频")
}
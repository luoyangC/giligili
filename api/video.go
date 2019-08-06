package api

import (
	"giligili/util"
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 添加视频
func CreateVideo(c *gin.Context) {
	claims := c.MustGet("claims").(*util.CustomClaims)
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(claims)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListVideo 获取视频列表
func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	res := service.List()
	c.JSON(200, res)
}

// ShowVideo 获取指定视频详情
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateVideo 更新指定视频
func UpdateVideo(c *gin.Context) {
	claims := c.MustGet("claims").(*util.CustomClaims)
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"), claims)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除指定视频
func DeleteVideo(c *gin.Context) {
	claims := c.MustGet("claims").(*util.CustomClaims)
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"), claims)
	c.JSON(200, res)
}

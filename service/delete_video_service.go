package service

import (
	"giligili/util"
	"giligili/model"
	"giligili/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct{}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string, c *util.CustomClaims) serializer.Response  {
	var video model.Video
	if err := model.DB.Where("user_id = ?", c.ID).First(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 4004,
			Msg: "没有找到该视频",
			Error: err.Error(),
		}
	}
	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.Response{
			Status: 5000,
			Msg: "删除视频失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 2000,
		Msg: "删除成功",
	}
}

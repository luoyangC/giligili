package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListVideoService 获取视频列表的服务
type ListVideoService struct{}

// List 获取视频列表
func (service *ListVideoService) List() serializer.Response  {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 5000,
			Msg: "数据库查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 2000,
		Data: serializer.BuildVideos(videos),
	}
}
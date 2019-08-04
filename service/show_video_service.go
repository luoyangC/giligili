package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 获取指定视频详情的服务
type ShowVideoService struct {}

// Show 获取指定视频详情
func (service *ShowVideoService) Show(id string) serializer.Response  {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 4004,
			Msg: "没有找到该视频",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 2000,
		Data: serializer.BuildVideo(&video),
	}
}
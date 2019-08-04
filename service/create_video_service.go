package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 创建视频的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=200"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info: service.Info,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.Response{
			Status: 5001,
			Msg: "视频保存错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(&video),
	}
}

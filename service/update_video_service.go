package service

import (
	"giligili/model"
	"giligili/serializer"
)

// UpdateVideoService 更新指定视频的服务
type UpdateVideoService struct{
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=200"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response  {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 4004,
			Msg: "没有找到该视频",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.Response{
			Status: 5000,
			Msg: "数据库保存出错",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 2000,
		Data: serializer.BuildVideo(&video),
	}
}
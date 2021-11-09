package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 视频投稿的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `form:"info" json:"info" binding:"max=3000"`
}

// Create 创建视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}

package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// 创建投稿视频
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

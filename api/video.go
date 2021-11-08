package api

import (
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

// 创建投稿视频
func CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "创建成功",
	})
}

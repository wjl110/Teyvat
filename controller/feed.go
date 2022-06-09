package controller

import (
	"douying/auth"
	"douying/common"
	"douying/dao"
	"douying/logic"
	"douying/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.VideoInfo `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}
func Feed(context *gin.Context)  {
	// 如果是登录状态则token不为空
	token := context.Query("token")
	username, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
	}
	var user models.User
	if token != "" {
		dao.RedisGet(username, &user)
	}

	context.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: logic.GetFeed(user, 0),
		NextTime:  time.Now().Unix(),
	})
}
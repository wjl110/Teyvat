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
	"strconv"
)


// 点赞
func FavoriteAction(context *gin.Context) {
	token := context.Query("token")
	username, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
		return
	}
	var user models.User
	if token != "" {
		dao.RedisGet(username, &user)
	}

	if user.Id != 0{
		videoId,_ := strconv.ParseInt(context.Query("video_id"), 10, 64)
		actionType,_ := strconv.ParseInt(context.Query("action_type"), 10, 64)
		favorite := models.UserFavorite{VideoId: videoId, UserId: user.Id, Status: actionType}
		models.SaveFavorite(&favorite)
		context.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		context.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// 点赞列表
func FavoriteList(context *gin.Context) {
	token := context.Query("token")
	_, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
		return
	}

	userId,_ := strconv.ParseInt(context.Query("user_id"), 10, 64)
	context.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: logic.FavoriteVideoFeed(userId),
	})
}
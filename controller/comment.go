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

type CommentListResponse struct {
	common.Response
	CommentList []logic.CommentInfo `json:"comment_list,omitempty"`
}

func CommentList(context *gin.Context)  {
	videoId,_ := strconv.ParseInt(context.Query("video_id"), 10, 64)
	context.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: logic.GetCommentListByVideo(videoId),
	})
}

func CommentAction(context *gin.Context)  {
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

	if user.Id != 0 {
		videoId,_ := strconv.ParseInt(context.Query("video_id"), 10, 64)
		actionType,_ := strconv.ParseInt(context.Query("action_type"), 10, 64)
		commentText := context.Query("comment_text")
		commentId,_ := strconv.ParseInt(context.Query("comment_id"), 10, 64)
		if actionType == 1 {
			models.AddComment(&models.Comment{UserId: user.Id, VideoId: videoId, Content: commentText,Status:1})
		} else {
			models.DeleteComment(&models.Comment{ID:commentId,Status:2})
		}

		context.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		context.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
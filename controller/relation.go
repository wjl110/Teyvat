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

type UserListResponse struct {
	common.Response
	UserList []common.UserInfo `json:"user_list"`
}
// 获取关注列表
func FollowList(context *gin.Context)  {
	token := context.Query("token")
	_, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
		return
	}
	userId,_ := strconv.ParseInt(context.Query("user_id"), 10, 64)

	context.JSON(http.StatusOK, UserListResponse{
		Response:  common.Response{StatusCode: 0},
		UserList: logic.GetUserFollowList(userId),
	})
}
//获取粉丝列表
func FollowerList(context *gin.Context)  {
	token := context.Query("token")
	_, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
		return
	}
	userId,_ := strconv.ParseInt(context.Query("user_id"), 10, 64)

	context.JSON(http.StatusOK, UserListResponse{
		Response:  common.Response{StatusCode: 0},
		UserList: logic.GetUserFollowerList(userId),
	})
}

// 关注操作
func RelationAction(context *gin.Context) {
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

	toUserId,_ := strconv.ParseInt(context.Query("to_user_id"), 10, 64)
	actionType,_ := strconv.ParseInt(context.Query("action_type"), 10, 64)

	if actionType == 1 {
	   models.AddFollow(&models.UserFollower{UserId: toUserId, FollowerId: user.Id})
	} else {
		models.DeleteFollow(toUserId, user.Id)
	}
	context.JSON(http.StatusOK, common.Response{StatusCode: 0})
}
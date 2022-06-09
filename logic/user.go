package logic

import (
	"douying/common"
	"douying/models"
)

func GetUserInfoByUserId(authorId int64,userId int64) common.UserInfo {
	// 获取作者用户信息
	user := models.User{
		Id: authorId,
	}
	models.GetUserById(&user)
	u := common.UserInfo{
		User: user,
		FollowCount: models.UserFollowCount(authorId),
		FollowerCount: models.UserFollowerCount(authorId),
		IsFollow: models.UserIsFollow(authorId, userId),
	}
	return u
}
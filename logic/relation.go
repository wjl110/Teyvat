package logic

import (
	"douying/common"
	"douying/models"
)

func GetUserFollowerList(userId int64) []common.UserInfo {
	var userFollowers []models.UserFollower
	models.GetUserFollower(userId, &userFollowers)
	list := make([]common.UserInfo,0)
	for _,v := range userFollowers {
		list = append(list, GetUserInfoByUserId(v.FollowerId, 0))
	}
	return list
}

func GetUserFollowList(userId int64) []common.UserInfo {
	var userFollowers []models.UserFollower
	models.GetUserFollow(userId, &userFollowers)
	list := make([]common.UserInfo,0)
	for _,v := range userFollowers {
		list = append(list, GetUserInfoByUserId(v.UserId, 0))
	}
	return list
}

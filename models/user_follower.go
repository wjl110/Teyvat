package models

import "douying/dao"

type UserFollower struct {
	Id int64
	UserId int64
	FollowerId int64
}
// 获取用户粉丝
func GetUserFollower(userId int64,userFollowers *[]UserFollower)  {
	dao.DB.Model(&UserFollower{}).Where("user_id = ?", userId).Find(userFollowers)
}
// 获取用户关注列表
func GetUserFollow(userId int64,userFollowers *[]UserFollower)  {
	dao.DB.Model(&UserFollower{}).Where("follower_id = ?", userId).Find(userFollowers)
}

// 用户粉丝数
func UserFollowerCount(userId int64) int64{
	var count int64
	dao.DB.Model(&UserFollower{}).Where("user_id = ?", userId).Count(&count)
	return count
}
// 用户关注数
func UserFollowCount(followerId int64) int64{
	var count int64
	dao.DB.Model(&UserFollower{}).Where("follower_id = ?", followerId).Count(&count)
	return count
}
// 判断用户之间是否关注
func UserIsFollow(userId int64, followerId int64) bool{
	var count int64
	dao.DB.Model(&UserFollower{}).Where("user_id = ? and follower_id = ?", userId, followerId).Count(&count)
	return count != 0
}
// 关注
func AddFollow(userFollower *UserFollower)  {
	dao.DB.Model(&UserFollower{}).Save(userFollower)
}
// 取消关注
func DeleteFollow(userId int64, followerId int64)  {
	dao.DB.Where("user_id = ? and follower_id = ?", userId, followerId).Delete(&UserFollower{})
}
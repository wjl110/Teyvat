package models

import (
	"douying/dao"
)

type UserFavorite struct {
	Id int64
	UserId int64
	VideoId int64
	Status int64
}

func IsFavorite(videoId int64, userId int64) bool {
	var count int64
	dao.DB.Model(&UserFavorite{}).Where("video_id = ? and user_id = ? and status = ?", videoId, userId, 1).Count(&count)
	return count != 0
}

func VideoFavoriteCount(videoId int64) int64{
	var count int64
	dao.DB.Model(&UserFavorite{}).Where("video_id = ? and status = ?", videoId, 1).Count(&count)
	return count
}

func SaveFavorite(favorite *UserFavorite)  {
	tx := dao.DB.Model(&UserFavorite{}).Where("id = ? and user_id = ?", favorite.VideoId, favorite.UserId).Update("status", favorite.Status)
	if tx.RowsAffected == 0 {
		dao.DB.Save(favorite)
	}
}

func FavoriteList(userId int64,favorite *[]UserFavorite)  {
	dao.DB.Model(&UserFavorite{}).Select("video_id","user_id").Where("user_id = ? and status = ?", userId, 1).Find(&favorite)
}

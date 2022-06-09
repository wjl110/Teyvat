package logic

import (
	"douying/common"
	"douying/models"
)

func FavoriteVideoFeed(userId int64)  []common.VideoInfo {
	var favorites []models.UserFavorite
	models.FavoriteList(userId, &favorites)
	var videos []models.Video
	videoIds := make([]int64,0)
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoId)
	}
	models.GetVideoInVideo(&videos, videoIds)
	videoInfos := make([]common.VideoInfo, 0)
	for _, video := range videos{
		videoInfo := common.VideoInfo{ Video: video,
			FavoriteCount: models.VideoFavoriteCount(video.Id),
			Author:        GetUserInfoByUserId(video.UserId, userId)}
			videoInfo.IsFavorite = models.IsFavorite(video.Id, userId)
		videoInfos = append(videoInfos, videoInfo)
	}
	return videoInfos
}

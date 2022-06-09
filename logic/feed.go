package logic

import (
	"douying/common"
	"douying/models"
)


func GetFeed(user models.User, userId int64)  []common.VideoInfo {
	var videos []models.Video
	// 判断需要限定条件获取对应用户视频
	if userId != 0 {
		models.GetVideosByUserId(&videos, userId)
	} else {
		models.GetAllVideos(&videos)
	}
	videoInfos := make([]common.VideoInfo, 0)
	for _, video := range videos{
		videoInfo := common.VideoInfo{ Video: video,
			FavoriteCount: models.VideoFavoriteCount(video.Id),
			CommentCount:  models.GetCommentCount(video.Id),
			Author:        GetUserInfoByUserId(video.UserId, user.Id)}
		if user.Id != 0 {
			videoInfo.IsFavorite = models.IsFavorite(video.Id, user.Id)
		}
		videoInfos = append(videoInfos, videoInfo)
	}

	return videoInfos
}

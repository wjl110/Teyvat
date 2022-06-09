package models

import (
	"douying/dao"
	"time"
)

type Video struct {
	Id int64  `json:"id,omitempty"`
	UserId int64 `json:"-"`
	PlayUrl string  `json:"play_url,omitempty"`
	CoverUrl string  `json:"cover_url,omitempty"`
	Title string  `json:"title,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
}

func GetAllVideos(video *[]Video) {
	dao.DB.Find(&video)
}

func GetVideosByUserId(video *[]Video, userId int64) {
	dao.DB.Select("id","user_id","play_url","cover_url","title","created_at").Where("user_id = ?", userId).Find(&video)
}
func GetVideoInVideo(video *[]Video,videoIds []int64)  {
	dao.DB.Where("id in ?", videoIds).Find(&video)
}

func SaveVideo(video *Video)  {
	dao.DB.Save(video)
}
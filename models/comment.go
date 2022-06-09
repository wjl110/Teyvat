package models

import (
	"douying/dao"
	"time"
)

type Comment struct {
	ID int64 `json:"id"`
	UserId int64 `json:"-"`
	VideoId int64 `json:"-"`
	Content string `json:"content"`
	CreatedAt time.Time  `json:"create_date,omitempty"`
	Status int `json:"-"`
}

func GetCommentListByVideo(videoId int64, comments *[]Comment)  {
	dao.DB.Model(&Comment{}).Where("video_id = ? and status = 1", videoId).Find(comments)
}

func GetCommentCount(videoId int64) int64 {
	var count int64
	dao.DB.Model(&Comment{}).Where("video_id = ? and status = ?", videoId, 1).Count(&count)
	return count
}

func AddComment(comment *Comment)  {
	dao.DB.Model(&Comment{}).Save(comment)
}

func DeleteComment(comment *Comment)  {
	dao.DB.Model(&Comment{}).Where("id = ?", comment.ID).Update("status", comment.Status)
}


package logic

import (
	"douying/common"
	"douying/models"
)

type CommentInfo struct {
	models.Comment `json:",inline"`
	User common.UserInfo `json:"user"`
}

func GetCommentListByVideo(videoId int64) []CommentInfo {
	var comments []models.Comment
	models.GetCommentListByVideo(videoId, &comments)
	commentInfos := make([]CommentInfo, 0)
	for _, comment := range comments {
		commentInfos = append(commentInfos, CommentInfo{comment, GetUserInfoByUserId(comment.UserId, 0)})
	}
	return commentInfos
}


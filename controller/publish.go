package controller

import (
	"douying/auth"
	"douying/common"
	"douying/logic"
	"douying/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.VideoInfo `json:"video_list"`
}
func PublishList(context *gin.Context) {
	token := context.Query("token")
	_, err := auth.CheckToken(token)
	if err != nil {
		fmt.Println("用户token错误")
		return
	}

	userId, _ := strconv.ParseInt(context.Query("user_id"), 10 , 64)
	context.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: logic.GetFeed(models.User{}, userId),
	})
}

func Publish(context *gin.Context) {
	userId, _ := strconv.ParseInt(context.PostForm("user_id"), 10 , 64)
	//username, err := auth.CheckToken(token)
	//if err != nil {
	//	fmt.Printf("用户token错误 %s\n", err)
	//	return
	//}
	//var user models.User
	//if token != "" {
	//	dao.RedisGet(username, &user)
	//}

	data, err := context.FormFile("data")
	if err != nil {
		context.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	//finalName := fmt.Sprintf("%s.mp3",  filename)
	if err := common.UploadOss(data, filename); err != nil {
		fmt.Printf("上传失败 %s\n", err)
		context.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	url := fmt.Sprintf("https://dou-yin-zzz.oss-cn-guangzhou.aliyuncs.com/%s", filename)
	coverUrl := fmt.Sprintf("%s?spm=qipa250&x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast", url)
	models.SaveVideo(&models.Video{UserId: userId, PlayUrl: url, CoverUrl: coverUrl, Title: filename})
	context.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  filename + " uploaded successfully",
	})
}

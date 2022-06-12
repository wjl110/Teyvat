package controller

import (
	"douying/auth"
	"douying/common"
	"douying/dao"
	"douying/logic"
	"douying/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}


type UserResponse struct {
	common.Response
	UserInfo common.UserInfo `json:"user"`
}


func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 通过用户名获取信息，判断是否重复
	user := models.User{Username: username}
	models.GetUserByUsername(&user)

	// id 默认为0 ，不为0则代表有数据，用户名重复
	if user.Id != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		u := models.User{Username: username, Password: password}
		models.SaveUser(&u)
		user.Id = u.Id
		token,_ := auth.CreateToken(username)
		err := dao.RedisSet(username, user)
		if err != nil {
			fmt.Printf("redis set Fail %v/n",err)
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   u.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := models.User{Username: username}
	models.GetUserByUsername(&user)
	if user.Password == password {
		token,_ := auth.CreateToken(username)
		err := dao.RedisSet(username, user)
		if err != nil {
			fmt.Printf("redis set Fail %v/n",err)
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	userInfo := logic.GetUserInfoByUserId(userId, 0)
	if userInfo.Username != "" {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			UserInfo: userInfo,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't login"},
		})
	}
}

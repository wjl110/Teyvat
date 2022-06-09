package models

import "douying/dao"

type User struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"name,omitempty"`
	Password string `json:"-"`
}

func GetUserById(user *User)  {
	dao.DB.Select("id","username").Find(&user, user.Id)
}

func GetUserByUsername(user *User)  {
	dao.DB.Where("username = ?", user.Username).First(&user)
}

func SaveUser(user *User)  {
	dao.DB.Save(user)
}


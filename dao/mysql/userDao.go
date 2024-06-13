package mysql

import (
	"bluebell/api"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"bluebell/utils/encrypted"
	"fmt"
)

type UserDao struct {
}

// UserExist  根据用户名查询用户
// 返回true表示用户不存在
func (u *UserDao) UserExist(username string) bool {
	var user models.User
	var count int64
	db.Where("user_name = ?", username).First(&user).Count(&count)
	if count == 0 {
		return true
	}
	return false
}

// InsertUser 插入用户
func (u *UserDao) InsertUser(data *models.UserSignUpParams) (err error) {
	var user models.User
	// 1.生成UID
	user_id := snowflake.GenID()
	user.User_id = user_id
	// 2.密码加密
	b_pwd, err := encrypted.HashPassword(data.Password)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = b_pwd
	user.User_name = data.User_name
	if err = db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Login 登录
// 返回err
func (u *UserDao) Login(user *models.User) (code int) {
	// 1.查询用户
	var count int64
	pwd := user.Password
	db.Where("user_name = ?", user.User_name).First(&user).Count(&count)
	if count == 0 {
		return api.CodeUserNotExist
	}
	// 2.校验密码
	if !encrypted.CheckPasswordHash(pwd, user.Password) {
		return api.CodeInvalidPassword
	}
	return api.CodeSuccess
}

func (u *UserDao) GetUserByID(id int64) (user *models.User, err error) {
	user = &models.User{}
	if err = db.Where("user_id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

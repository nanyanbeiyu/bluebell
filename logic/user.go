package logic

import (
	"bluebell/api"
	"bluebell/dao/mysql"
	"bluebell/middleware/jwt"
	"bluebell/models"
)

var userDao *mysql.UserDao

type UserLg struct {
}

func (u *UserLg) SignUp(data *models.UserSignUpParams) (code int) {
	// 1.判断用户是否存在
	exist := userDao.UserExist(data.User_name)
	if !exist {
		return api.CodeUserExist
	}
	// 2.保存进数据库
	if err := userDao.InsertUser(data); err != nil {
		return api.CodeServerBusy
	}
	return api.CodeSuccess
}

func (u *UserLg) Login(data *models.UserLoginParams) (code int, token string) {
	var user *models.User
	user = &models.User{
		User_name: data.User_name,
		Password:  data.Password,
	}
	// check
	code = userDao.Login(user)
	if code != api.CodeSuccess {
		return code, ""
	}

	token, _ = jwt.CreateToken(user.User_id, user.User_name)
	return code, token
}

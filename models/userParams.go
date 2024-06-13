package models

// UserSignUpParams 用户注册参数
type UserSignUpParams struct {
	User_name   string `json:"user_name" binding:"required"`                    // 用户名 必须传入
	Password    string `json:"password" binding:"required"`                     // 密码 必须传入
	Re_password string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码 必须传入
}

// UserLoginParams 用户登录参数
type UserLoginParams struct {
	User_name string `json:"user_name" binding:"required"` // 用户名 必须传入
	Password  string `json:"password" binding:"required"`  // 密码 必须传入
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // 默认字段
	User_id    int64  `gorm:"not null;unique"` // 用户id
	User_name  string `gorm:"not null;unique"` // 用户名
	Password   string `gorm:"not null"`        // 密码
	Email      string // 邮箱
	Gender     int8   `gorm:"not null;default:0"` // 性别
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "User"
}

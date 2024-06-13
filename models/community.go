package models

import "gorm.io/gorm"

type Community struct {
	gorm.Model            // 默认字段
	Community_id   int64  `gorm:"not null;unique"`                                // 社区id
	Community_name string `gorm:"not null;unique"`                                // 社区名称
	Introduction   string `gorm:"not null"`                                       // 社区介绍
	Posts          []Post `gorm:"foreignKey:CommunityID;references:Community_id"` // 社区对应的帖子
}

func (c *Community) TableName() string {
	return "community"
}

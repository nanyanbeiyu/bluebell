package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model         // 默认字段
	Title       string `gorm:"not null"`           // 帖子标题
	Content     string `gorm:"not null;type:text"` // 帖子内容
	AuthorID    int64  `gorm:"not null"`           // 帖子作者的ID
	CommunityID int64  `gorm:"not null"`           // 帖子所属社区
	PostID      int64  `gorm:"not null;unique"`    // 帖子ID
	Status      int8   `gorm:"not null;default:0"` // 帖子状态
}

func (p *Post) TableName() string {
	return "post"
}

type ApiPostDetail struct {
	AuthorName     string        `json:"author_name"`    // 帖子作者名称
	Community_name string        `json:"community_name"` // 帖子所属社区名称
	*Post          `json:"post"` // 帖子信息
}

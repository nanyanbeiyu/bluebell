package models

// CaretCommunityParams 创建社区参数
type CaretCommunityParams struct {
	Community_name string `json:"community_name" binding:"required"` // 社区名称
	Introduction   string `json:"introduction" binding:"required"`   // 社区介绍
}

// GetCommunityListParams 获取社区列表参数
type GetCommunityListParams struct {
	Community_name string `json:"community_name"` // 社区名称
	Introduction   string `json:"introduction"`   // 社区介绍
	Community_id   int64  `json:"community_id"`   // 社区ID
}

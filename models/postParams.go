package models

// CartePostParams 创建帖子参数
type CartePostParams struct {
	Title       string `json:"title" binding:"required"`               // 帖子标题
	Content     string `json:"content" binding:"required"`             // 帖子内容
	CommunityID int64  `json:"community_id,string" binding:"required"` // 社区ID
}

// PostListParams 帖子列表参数
type PostListParams struct {
	CommunityID int64  `json:"community_id" form:"community_id"` // 社区ID
	Page        int64  `json:"page" form:"page"`                 // 页码
	Size        int64  `json:"size" form:"size"`                 // 每页数据量
	Order       string `json:"order" form:"order"`               // 排序
}

// CmPostListParams 社区帖子列表参数
type CmPostListParams struct {
	PostListParams // 嵌入帖子列表参数
}

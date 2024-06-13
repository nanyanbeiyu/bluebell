package models

type VoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`       // 帖子ID
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票（1）反对票（-1）取消投票（0）
}

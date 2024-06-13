package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能：
// 投一票加432分 86400秒/200票 -》 200张赞成票可以续一天时间
/*
投票的业务逻辑：
	1. 投票数据要存储到redis中
	2. 投票的时候，需要向redis发送一条消息
	3. redis中的消息，有worker去消费，然后修改数据库中的数据
*/
/* 投票的几种情况：
Direction = 1:
	1. 之前没投过票, 投赞成票
	2. 之前投反对票，修改为投赞成票
Direction = 0:
	1. 之前没投过票，啥也不做
	2. 之前投赞成票，取消投票
	3. 之前投反对票，取消投票
Direction = -1:
	1. 之前投过赞成票，修改为投反对票
	2. 之前没投过票, 投反对票

投票的限制：
	自帖子发布之日起一个星期内允许用户投票，超过一周后不允许投票
		1. 到期之后将redis 中的投票数据同步到数据库中
		2.到期之后删除redis中的投票数据
*/

type VoteLg struct {
}

// VoteForPost 帖子投票函数
func (v *VoteLg) VoteForPost(uid int64, data *models.VoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", uid),
		zap.Int64("PostID", data.PostID),
		zap.Int8("direction", data.Direction))
	return redis.VoteForPost(strconv.Itoa(int(uid)), strconv.Itoa(int(data.PostID)), float64(data.Direction))
}

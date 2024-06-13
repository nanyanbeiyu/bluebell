package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	socrePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过期")
	ErrVoteRepeat     = errors.New("不允许重复投票")
)

// CreatePost 创建帖子
func CreatePost(postID, communityID int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 把帖子社区id添加到社区set
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(context.Background(), ckey)
	_, err := pipeline.Exec(context.Background())

	return err
}

// VoteForPost 投票
func VoteForPost(uid, postID string, direction float64) error {
	// 1.判断投票限制
	postTime := rdb.ZScore(context.Background(), getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2.更新分数
	// 先查当前帖子的投票记录
	ov := rdb.ZScore(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), uid).Val()
	if ov == direction {
		return ErrVoteRepeat
	}
	var dir float64
	if direction > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - direction)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZSet), dir*diff*socrePerVote, postID)

	// 3.记录用户为该帖子投票的数据
	if direction == 0 {
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), uid)

	} else {
		pipeline.ZAdd(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), &redis.Z{
			Score:  direction,
			Member: uid,
		})
	}
	_, err := pipeline.Exec(context.Background())
	return err
}

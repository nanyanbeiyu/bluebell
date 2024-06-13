package redis

import (
	"bluebell/models"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func getIDSFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3.ZRevRange 按分数从高到低的顺序查询指定数量圆度
	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}

// GetPostIDByOrder 根据给定的排序顺序从redis中获取帖子的id
func GetPostIDByOrder(p *models.PostListParams) ([]string, error) {
	// 1.从redis获取id
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询的索引起始点
	return getIDSFormKey(key, p.Page, p.Size)
}

func GetCommunityPostIDsInOrder(p *models.PostListParams) (data []string, err error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// ZInterStore 把分区的帖子set和分数的zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑获取数据
	// 社区的key
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key减少ZInterStore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(context.Background(), key).Val() < 1 {
		// 不存在 需要计算
		pipeline := redis.Pipeline{}
		pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Keys:      []string{ckey, orderKey},
			Aggregate: "MAX",
		})
		pipeline.Expire(context.Background(), key, 60*time.Second) // 设置过期时间
		_, err = pipeline.Exec(context.Background())
		if err != nil {
			return nil, err
		}

	}
	return getIDSFormKey(key, p.Page, p.Size)
}

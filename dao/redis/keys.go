package redis

const (
	KeyPrefix          = "bluebell:"   // redis key 前缀
	KeyPostTimeZSet    = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset; 帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // zset; 记录用户及投票类型; 参数是post id
	KeyCommunitySetPF  = "community:"  // set; 记录每个分区下帖子的 id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}

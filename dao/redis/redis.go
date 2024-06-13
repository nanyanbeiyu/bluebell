package redis

import (
	"bluebell/settings"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(conf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Post),
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
		PoolSize: conf.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return nil
}

func Close() {
	_ = rdb.Close()
}

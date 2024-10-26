package infrastructure

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func InitializeRedSyncConn(redisClient *redis.Client) *redsync.Redsync {
	redisPool := goredis.NewPool(redisClient)
	rs := redsync.New(redisPool)

	return rs
}

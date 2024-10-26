package infrastructure

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitializeRedisConn() *redis.Client {
	logger := logrus.WithField("cacheHost", config.Env.Redis.CacheHost)

	opts, err := redis.ParseURL(config.Env.Redis.CacheHost)
	if err != nil {
		logger.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         opts.Addr,
		Username:     opts.Username,
		Password:     opts.Password,
		DB:           opts.DB,
		DialTimeout:  config.Env.Redis.DialTimeout,
		WriteTimeout: config.Env.Redis.WriteTimeout,
		ReadTimeout:  config.Env.Redis.ReadTimeout,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal("err connect to Redis: ", err)
	}

	return rdb
}

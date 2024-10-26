package infrastructure

import (
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitializeAsynqServer() *asynq.Server {
	redisConn := GetAsynqRedisConn()
	server := asynq.NewServer(
		redisConn,
		asynq.Config{
			Concurrency: config.Env.Asynq.WorkerConcurrency,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return server
}

func InitializeAsynqClient() *asynq.Client {
	redisConn := GetAsynqRedisConn()
	client := asynq.NewClient(redisConn)

	return client
}

func InitializeAsynqScheduler() *asynq.Scheduler {
	redisConn := GetAsynqRedisConn()
	scheduler := asynq.NewScheduler(
		redisConn,
		&asynq.SchedulerOpts{},
	)

	return scheduler
}

func GetAsynqRedisConn() asynq.RedisConnOpt {

	opts, err := redis.ParseURL(config.Env.Redis.WorkerCacheHost)
	if err != nil {
		logrus.Fatal(err)
	}

	redisConn := asynq.RedisClientOpt{
		Addr:         opts.Addr,
		Username:     opts.Username,
		Password:     opts.Password,
		DB:           opts.DB,
		DialTimeout:  config.Env.Redis.DialTimeout,
		WriteTimeout: config.Env.Redis.WriteTimeout,
		ReadTimeout:  config.Env.Redis.ReadTimeout,
	}

	return redisConn
}

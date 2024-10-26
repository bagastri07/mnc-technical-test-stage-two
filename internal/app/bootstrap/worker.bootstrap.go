package bootstrap

import (
	"context"
	"os"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/infrastructure"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/distributed"
	userHdr "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/distributed/user"
	userRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/user"
	userBalanceRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/userbalance"
	userTransactionRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/usertransaction"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/usertransferhistory"
	userUC "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/usecase/user"
	"github.com/go-redsync/redsync/v4"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppWorker struct {
	PostgresDB     *gorm.DB
	RedisClient    *redis.Client
	AysnqServer    *asynq.Server
	AsynqScheduler *asynq.Scheduler
	AsynqClient    *asynq.Client
	RedSync        *redsync.Redsync
}

func greetingWorker() {
	logrus.Info(config.Env.App.Name, " Worker Started with PID: ", os.Getpid())
}

func StartWorker() {
	greetingWorker()

	redisClient := infrastructure.InitializeRedisConn()
	app := AppWorker{
		PostgresDB:     infrastructure.InitializePostgresConn(),
		RedisClient:    redisClient,
		AysnqServer:    infrastructure.InitializeAsynqServer(),
		AsynqScheduler: infrastructure.InitializeAsynqScheduler(),
		AsynqClient:    infrastructure.InitializeAsynqClient(),
		RedSync:        infrastructure.InitializeRedSyncConn(redisClient),
	}

	pgDB, err := app.PostgresDB.DB()
	util.ContinueOrFatal(err)

	// init repositories
	userRepository := userRepo.New().
		WithPostgresqlDB(app.PostgresDB)
	userBalanceRepository := userBalanceRepo.New().
		WithPostgresqlDB(app.PostgresDB)
	userTransactionRepository := userTransactionRepo.New().
		WithPostgresqlDB(app.PostgresDB)
	userTransferHistoryRepository := usertransferhistory.New().
		WithPostgresqlDB(app.PostgresDB)

	// init useCases
	userUseCase := userUC.New().
		WithPostgresqlDB(app.PostgresDB).
		WithRedSync(app.RedSync).
		WithUserRepository(userRepository).
		WithUserBalanceRepository(userBalanceRepository).
		WithUserTransactionRepository(userTransactionRepository).
		WithUserTransferHistoryRepository(userTransferHistoryRepository)

	// init handlers
	userHandler := userHdr.New().
		WithUserUseCase(userUseCase)

	mux := distributed.NewServeMuxBuilder().
		WithUserHandler(userHandler).
		Build()

	go func() {
		if err := app.AysnqServer.Run(mux); err != nil {
			logrus.Fatal(err)
		}
	}()

	wait := util.GracefulShutdown(context.Background(), config.Env.App.GracefulShutdownTimeOut, map[string]util.Operation{
		"postgressql connection": func(_ context.Context) error {
			return pgDB.Close()
		},
		"asynq connection": func(_ context.Context) error {
			app.AysnqServer.Stop()
			return nil
		},
		"asynq client": func(_ context.Context) error {
			return app.AsynqClient.Close()
		},
		"asynq scheduler connection": func(_ context.Context) error {
			app.AsynqScheduler.Shutdown()
			return nil
		},
		"redis connection": func(_ context.Context) error {
			return app.RedisClient.Close()
		},
	})

	<-wait
}

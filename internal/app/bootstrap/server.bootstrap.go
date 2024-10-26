package bootstrap

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/infrastructure"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest"
	healthCheckHdr "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/healthcheck"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/middleware"
	userHdr "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/user"
	userRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/user"
	userBalanceRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/userbalance"
	userTransactionRepo "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/repository/usertransaction"
	userUC "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/usecase/user"
	"github.com/common-nighthawk/go-figure"
	"github.com/go-gomail/gomail"
	"github.com/go-redsync/redsync/v4"
	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type appServer struct {
	PostgresDB    *gorm.DB
	RedisClient   *redis.Client
	RedSync       *redsync.Redsync
	HTPPGinServer *http.Server
	AsynqClient   *asynq.Client
	GomailDialer  *gomail.Dialer
}

func greetingServer() {
	figure.NewColorFigure(config.Env.App.Name, "thin", "blue", true).Print()
	logrus.Info(config.Env.App.Quote)
}

func StartServer() {
	greetingServer()

	redisClient := infrastructure.InitializeRedisConn()
	app := appServer{
		HTPPGinServer: infrastructure.InitializeGinServer(),
		PostgresDB:    infrastructure.InitializePostgresConn(),
		RedisClient:   redisClient,
		AsynqClient:   infrastructure.InitializeAsynqClient(),
		RedSync:       infrastructure.InitializeRedSyncConn(redisClient),
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

	// init useCases
	userUseCase := userUC.New().
		WithPostgresqlDB(app.PostgresDB).
		WithAsynqClient(app.AsynqClient).
		WithRedSync(app.RedSync).
		WithUserRepository(userRepository).
		WithUserBalanceRepository(userBalanceRepository).
		WithUserTransactionRepository(userTransactionRepository)

	// init handlers
	healthCheckHandler := healthCheckHdr.New()
	userHandler := userHdr.New().
		WithUserUseCase(userUseCase)

	// init middleware
	jetMid := middleware.NewJwtMiddleware()

	// init routes
	app.HTPPGinServer.Handler = rest.InitRoutes(
		jetMid,
		healthCheckHandler,
		userHandler,
	)

	go func() {
		err := app.HTPPGinServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			util.ContinueOrFatal(err)
		}
	}()

	go func() {
		if config.Env.Env != "production" {
			serveAsynqMonitoring()
		}
	}()

	wait := util.GracefulShutdown(context.Background(), config.Env.App.GracefulShutdownTimeOut, map[string]util.Operation{
		"postgres connection": func(ctx context.Context) error {
			return pgDB.Close()
		},
		"redis connection": func(ctx context.Context) error {
			return app.RedisClient.Close()
		},
		"http gin server": func(ctx context.Context) error {
			return app.HTPPGinServer.Close()
		},
		"asynq client": func(ctx context.Context) error {
			return app.AsynqClient.Close()
		},
	})
	<-wait
}

func serveAsynqMonitoring() {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: infrastructure.GetAsynqRedisConn(),
	})

	r := mux.NewRouter()
	r.PathPrefix(h.RootPath()).Handler(h)

	srv := &http.Server{
		Handler:           r,
		Addr:              ":4050",
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Go to http://localhost:4050/monitoring to see asynqmon homepage.
	log.Fatal(srv.ListenAndServe())
}

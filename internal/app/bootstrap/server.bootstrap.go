package bootstrap

import (
	"context"
	"errors"
	"net/http"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/bagastri07/gin-boilerplate-service/internal/infrastructure"
	"github.com/bagastri07/gin-boilerplate-service/internal/svc/handler/rest"
	healthCheckHdr "github.com/bagastri07/gin-boilerplate-service/internal/svc/handler/rest/healthcheck"
	"github.com/bagastri07/gin-boilerplate-service/internal/svc/handler/rest/middleware"
	userHdr "github.com/bagastri07/gin-boilerplate-service/internal/svc/handler/rest/user"
	userRepo "github.com/bagastri07/gin-boilerplate-service/internal/svc/repository/user"
	userUC "github.com/bagastri07/gin-boilerplate-service/internal/svc/usecase/user"
	"github.com/common-nighthawk/go-figure"
	"github.com/go-gomail/gomail"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type appServer struct {
	PostgresDB    *gorm.DB
	RedisClient   *redis.Client
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

	app := appServer{
		HTPPGinServer: infrastructure.InitializeGinServer(),
		PostgresDB:    infrastructure.InitializePostgresConn(),
		RedisClient:   infrastructure.InitializeRedisConn(),
		AsynqClient:   infrastructure.InitializeAsynqClient(),
	}

	pgDB, err := app.PostgresDB.DB()
	util.ContinueOrFatal(err)

	// init repositories
	userRepository := userRepo.New().
		WithPostgresqlDB(app.PostgresDB)

	// init useCases
	userUseCase := userUC.New().
		WithUserRepository(userRepository)

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

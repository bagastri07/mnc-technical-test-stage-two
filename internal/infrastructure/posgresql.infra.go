package infrastructure

import (
	"time"

	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/jpillora/backoff"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	postgesDB *gorm.DB
	// StopTickerCh signal for closing ticker channel
	StopTickerCh chan bool
)

func InitializePostgresConn() *gorm.DB {
	// Create the PostgreSQL DSN

	conn, err := openPostgresConn(config.Env.Postgres.DSN)
	if err != nil {
		log.WithField("databaseDSN", config.Env.Postgres.DSN).Fatal("failed to connect postgresql database: ", err)
	}

	StopTickerCh = make(chan bool)

	go checkConnection(time.NewTicker(config.Env.Postgres.PingInterval), config.Env.Postgres.DSN)

	switch config.Env.Postgres.LogLevel {
	case "error":
		conn.Logger = conn.Logger.LogMode(logger.Error)
	case "warn":
		conn.Logger = conn.Logger.LogMode(logger.Warn)
	case "silent":
		conn.Logger = conn.Logger.LogMode(logger.Silent)
	default:
		conn.Logger = conn.Logger.LogMode(logger.Info)
	}

	postgesDB = conn

	return postgesDB
}

func checkConnection(ticker *time.Ticker, dsn string) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := postgesDB.DB(); err != nil {
				reconnectPostgresConn(dsn)
			}
		}
	}
}

func reconnectPostgresConn(dsn string) {
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    1 * time.Second,
	}

	postgresRetryAttempts := config.Env.Postgres.RetryAttempts

	for b.Attempt() < postgresRetryAttempts {
		conn, err := openPostgresConn(dsn)
		if err != nil {
			log.WithField("databaseDSN", dsn).Error("failed to connect postgresql database: ", err)
		}

		if conn != nil {
			postgesDB = conn
			break
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= postgresRetryAttempts {
		log.Fatal("maximum retry to connect database")
	}
	b.Reset()
}

func openPostgresConn(dsn string) (*gorm.DB, error) {
	psqlDialector := postgres.Open(dsn)
	db, err := gorm.Open(psqlDialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	conn.SetMaxIdleConns(config.Env.Postgres.MaxIdleConns)
	conn.SetMaxOpenConns(config.Env.Postgres.MaxOpenConns)
	conn.SetConnMaxLifetime(config.Env.Postgres.ConnMaxLifetime)

	return db, nil
}

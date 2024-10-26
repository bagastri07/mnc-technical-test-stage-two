package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Env *EnvConfig
)

type EnvConfig struct {
	Env      string   `mapstructure:"env"`
	App      App      `mapstructure:"app"`
	JWT      JWT      `mapstructure:"jwt"`
	Crypto   Crypto   `mapstructure:"crypto"`
	Postgres Postgres `mapstructure:"postgres"`
	Redis    Redis    `mapstructure:"redis"`
	Asynq    Asynq    `mapstructure:"broker"`
	Aws      Aws      `mapstructure:"aws"`
	Sentry   Sentry   `mapstructure:"sentry"`
	Mailer   Mailer   `mapstructure:"mailer"`
	HTTP     HTTP     `mapstructure:"http"`
}

type App struct {
	Name                    string        `mapstructure:"name"`
	Quote                   string        `mapstructure:"quote"`
	LogLevel                string        `mapstructure:"log_level"`
	Port                    string        `mapstructure:"port"`
	GracefulShutdownTimeOut time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type JWT struct {
	UserSecret string        `mapstructure:"user_secret"`
	Timeout    time.Duration `mapstructure:"timeout"`
	MaxRefresh time.Duration `mapstructure:"max_refresh"`
}

type Crypto struct {
	Salt string `mapstructure:"salt"`
}

type Postgres struct {
	DSN             string        `mapstructure:"dsn"`
	LogLevel        string        `mapstructure:"log_level"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	PingInterval    time.Duration `mapstructure:"ping_interval"`
	RetryAttempts   float64       `mapstructure:"retry_attempts"`
}

type Redis struct {
	CacheHost       string        `mapstructure:"cache_host"`
	WorkerCacheHost string        `mapstructure:"worker_cache_host"`
	DialTimeout     time.Duration `mapstructure:"dial_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
}

type Asynq struct {
	WorkerConcurrency int           `mapstructure:"worker_concurrency"`
	Retention         time.Duration `mapstructure:"retention"`
	MaxRetry          int           `mapstructure:"max_retry"`
}

type Aws struct {
	Region           string        `mapstructure:"region"`
	AccessKey        string        `mapstructure:"access_key"`
	SecretKey        string        `mapstructure:"secret_key"`
	Bucket           string        `mapstructure:"bucket"`
	SignedExpiration time.Duration `mapstructure:"signed_expiration"`
}

type Sentry struct {
	DSN          string        `mapstructure:"dsn"`
	EnableAPM    bool          `mapstructure:"enable_apm"`
	SampleRate   float64       `mapstructure:"sample_rate"`
	Timeout      time.Duration `mapstructure:"timeout"`
	IgnoreErrors []string      `mapstructure:"ignore_errors"`
}

type Mailer struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type HTTP struct {
	Timeout              time.Duration `mapstructure:"timeout" default:"20s"`
	DisableKeepAlives    bool          `mapstructure:"disable_keep_alives"`
	MaxConsPerHost       int           `mapstructure:"max_cons_per_host"`
	MaxIdleConns         int           `mapstructure:"max_idle_cons"`
	MaxIdleConsPerHost   int           `mapstructure:"max_idle_cons_per_host"`
	IdleConnTimeout      time.Duration `mapstructure:"idle_conn_timeout" default:"60s"`
	TLSHandshakeTimeout  time.Duration `mapstructure:"tls_handshake_timeout" default:"10s"`
	DialContextTimeout   time.Duration `mapstructure:"dial_context_timeout"  default:"30s"`
	DialContextKeepAlive time.Duration `mapstructure:"dial_context_keep_alive"  default:"30s"`
}

func LoadConfig() {
	// Initialize viper
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Failed to read config file: ", err)
	}

	// Unmarshal the config file into the Config struct
	err = viper.Unmarshal(&Env)
	if err != nil {
		logrus.Fatal("Failed to unmarshal config file: ", err)
	}
}

package infrastructure

import (
	"net/http"
	"time"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/gin-gonic/gin"
)

func InitializeGinServer() *http.Server {
	switch config.Env.Env {
	case constant.EnvDevelopment, constant.EnvStaging:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	util.AddValidation()

	return &http.Server{
		Addr:           ":" + config.Env.App.Port,
		ReadTimeout:    time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
}

package infrastructure

import (
	"net/http"
	"time"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
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

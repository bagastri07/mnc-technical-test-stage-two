package rest

import (
	"net/http"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/healthcheck"
	jwtMidd "github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/middleware"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/rest/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	jwtMid *jwtMidd.JwtMiddleware,
	healthCheckHdr *healthcheck.Handler,
	userHandler *user.Handler,
) *gin.Engine {
	r := gin.New()

	// using cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           nil,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1Group := r.Group("/v1")
	v1Group.GET("/ping", healthCheckHdr.Ping)

	v1Group.POST("/login", userHandler.PostLogin)
	v1Group.POST("/register", userHandler.PostRegister)

	v1Group.Use(jwtMid.UserAuth())
	v1Group.POST("/topup", userHandler.PostTopUp)
	v1Group.POST("/pay", userHandler.PostPayment)
	v1Group.POST("/transfer", userHandler.PostTransfer)
	v1Group.GET("/transactions", userHandler.GetTransactions)
	v1Group.PUT("/profile", userHandler.PutProfile)

	return r
}

package middleware

import (
	"context"
	"strings"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/claim"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JwtMiddleware struct{}

func NewJwtMiddleware() *JwtMiddleware {
	return new(JwtMiddleware)
}

func (m *JwtMiddleware) UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.HandleError(c, constant.ErrUnauthorized)
			c.Abort()
			return
		}

		// Split the token from the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.HandleError(c, constant.ErrUnauthorized)
			c.Abort()
			return
		}

		userClaim := &claim.User{}
		userClaim, err := util.ValidateJWT[*claim.User](tokenString, []byte(config.Env.JWT.UserSecret), userClaim)
		if err != nil {
			logrus.Error(err)
			response.HandleError(c, err)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constant.IDKey, userClaim.ID)
		ctx = context.WithValue(ctx, constant.UsernameKey, userClaim.Username)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

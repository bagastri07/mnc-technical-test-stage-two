package user

import (
	"context"
	"time"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/claim"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) Login(ctx context.Context, req request.UserLoginReq) (*response.TokenResponse, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if ok := util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt) == user.Password; !ok {
		return nil, constant.ErrUnauthorized
	}

	expiredAt := time.Now().Add(config.Env.JWT.Timeout)
	token, err := util.GenerateToken[claim.User]([]byte(config.Env.JWT.UserSecret), claim.User{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Env.App.Name,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.TokenResponse{
		Token:     *token,
		ExpiredAt: expiredAt.Format(time.RFC3339),
	}, nil
}

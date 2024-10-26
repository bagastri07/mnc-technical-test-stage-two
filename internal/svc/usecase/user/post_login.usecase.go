package user

import (
	"context"
	"errors"
	"time"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/claim"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (u *UseCase) PostLogin(ctx context.Context, req request.PostUserLoginReq) (*response.PostLoginUserResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByPhoneNumber(ctx, req.PhoneNumber)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrPhoneNumberAndPINDoesntMatch
	}
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if ok := util.EncryptWithSalt(req.PIN, config.Env.Crypto.Salt) == user.PIN; !ok {
		return nil, constant.ErrPhoneNumberAndPINDoesntMatch
	}

	expiredAt := time.Now().Add(config.Env.JWT.Timeout)
	token, err := util.GenerateToken[claim.User]([]byte(config.Env.JWT.UserSecret), claim.User{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Env.App.Name,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	refreshExpiredAt := time.Now().Add(config.Env.JWT.RefreshTimeOut)
	refreshToken, err := util.GenerateToken[claim.User]([]byte(config.Env.JWT.UserRefreshSecret), claim.User{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Env.App.Name,
			ExpiresAt: jwt.NewNumericDate(refreshExpiredAt),
		},
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.PostLoginUserResp{
		AccessToken:  *token,
		RefreshToken: *refreshToken,
	}, nil
}

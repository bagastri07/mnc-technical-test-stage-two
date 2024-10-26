package user

import (
	"context"
	"errors"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/entity"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/guregu/null/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (u *UseCase) Register(ctx context.Context, req request.UserRegisterReq) (*response.UserIDResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err)
		return nil, err
	}

	if user != nil {
		return nil, constant.ErrUserAlreadyExist
	}

	hashedPassword := util.EncryptWithSalt(req.Password, config.Env.Crypto.Salt)

	newUser := &entity.User{
		Username: req.Username,
		Password: hashedPassword,
		FullName: null.StringFromPtr(req.FullName).NullString,
	}

	err = u.userRepository.Upsert(ctx, newUser)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.UserIDResp{UserID: newUser.ID}, nil
}

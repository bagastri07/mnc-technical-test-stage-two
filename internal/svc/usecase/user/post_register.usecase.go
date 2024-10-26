package user

import (
	"context"
	"errors"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/config"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (u *UseCase) PostRegister(ctx context.Context, req request.PostUserRegisterReq) (*response.PostUserRegisterResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
		tx  = u.db.WithContext(ctx).Begin()
		err error
	)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logger.Panic(p)
		}
		if err != nil {
			tx.Rollback()
			logger.Error(err)
		} else {
			tx.Commit()
		}
	}()

	user, err := u.userRepository.FindByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err)
		return nil, err
	}

	if user != nil {
		return nil, constant.ErrPhoneNumberAlreadyRegistered
	}

	hashedPIN := util.EncryptWithSalt(req.PIN, config.Env.Crypto.Salt)

	newUser := &entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		PIN:         hashedPIN,
		Address:     null.StringFromPtr(req.Address),
	}

	err = u.userRepository.Upsert(ctx, newUser)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	err = u.userBalanceRepository.Upsert(ctx, &entity.UserBalance{
		UserID:  newUser.ID,
		Balance: decimal.Zero,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.PostUserRegisterResp{
		UserID:      newUser.ID,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
		CreatedAt:   newUser.CreatedAt,
	}, nil
}

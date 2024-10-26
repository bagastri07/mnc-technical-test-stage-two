package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	var (
		user   = new(entity.User)
		logger = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("phone_number = ?", phoneNumber).
		First(user).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"phoneNumber": phoneNumber,
		}).Error(err)
		return nil, err
	}

	return user, nil
}

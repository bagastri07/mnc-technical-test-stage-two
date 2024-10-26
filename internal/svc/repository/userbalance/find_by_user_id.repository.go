package userbalance

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserBalance, error) {
	var (
		user   = new(entity.UserBalance)
		logger = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(user).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"userID": userID.String(),
		}).Error(err)
		return nil, err
	}

	return user, nil
}

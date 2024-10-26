package usertransaction

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *Repository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserTransaction, error) {
	var (
		userTransaction = make([]entity.UserTransaction, 0)
		logger          = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&userTransaction).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"ID": userID.String(),
		}).Error(err)
		return nil, err
	}

	return userTransaction, nil
}

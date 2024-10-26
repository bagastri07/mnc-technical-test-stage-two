package usertransaction

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.UserTransaction, error) {
	var (
		userTransaction = new(entity.UserTransaction)
		logger          = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		First(userTransaction).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"ID": ID.String(),
		}).Error(err)
		return nil, err
	}

	return userTransaction, nil
}

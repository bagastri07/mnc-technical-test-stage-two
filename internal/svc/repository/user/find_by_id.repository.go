package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.User, error) {
	var (
		user   = new(entity.User)
		logger = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		First(user).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"ID": ID.String(),
		}).Error(err)
		return nil, err
	}

	return user, nil
}

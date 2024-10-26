package user

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/entity"
	"github.com/sirupsen/logrus"
)

func (r *Repository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var (
		user   = new(entity.User)
		logger = util.GetLoggerFromCtx(ctx)
	)
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(user).
		Error

	if err != nil {
		logger.WithFields(logrus.Fields{
			"username": username,
		}).Error(err)
		return nil, err
	}

	return user, nil
}

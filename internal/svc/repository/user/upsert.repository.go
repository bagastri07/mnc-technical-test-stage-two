package user

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/entity"
	"github.com/sirupsen/logrus"
)

func (r *Repository) Upsert(ctx context.Context, user *entity.User) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Save(&user).
		Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user": util.Dump(user),
		}).Error(err)
	}

	return nil
}

package usertransaction

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/sirupsen/logrus"
)

func (r *Repository) Upsert(ctx context.Context, userTransaction *entity.UserTransaction) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Save(&userTransaction).
		Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user": util.Dump(userTransaction),
		}).Error(err)
		return err
	}

	return nil
}

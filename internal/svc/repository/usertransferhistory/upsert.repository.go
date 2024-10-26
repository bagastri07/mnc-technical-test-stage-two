package usertransferhistory

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/sirupsen/logrus"
)

func (r *Repository) Upsert(ctx context.Context, userTransferHistory *entity.UserTransferHistory) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Save(&userTransferHistory).
		Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user": util.Dump(userTransferHistory),
		}).Error(err)
		return err
	}

	return nil
}

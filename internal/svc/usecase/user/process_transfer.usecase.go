package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/cachekey"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/guregu/null/v5"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) ProcessTransfer(ctx context.Context, req request.ProcessTransferReq) error {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
		tx                      = u.db.WithContext(ctx).Begin()
		senderTransactionKey    = cachekey.NewUserTransactionKey(req.UserID)
		senderMutex             = util.RedSyncObtainLock(u.redSync, senderTransactionKey)
		recipientTransactionKey = cachekey.NewUserTransactionKey(req.TargetUserID)
		recipientMutex          = util.RedSyncObtainLock(u.redSync, recipientTransactionKey)
		err                     error
	)

	util.RedSyncLock(ctx, recipientMutex)
	defer util.RedSyncUnlock(ctx, recipientMutex)
	util.RedSyncLock(ctx, senderMutex)
	defer util.RedSyncUnlock(ctx, senderMutex)

	senderTransaction, err := u.userTransactionRepository.FindByID(ctx, req.SenderTransactionID)
	if err != nil {
		logger.Error(err)
		return err
	}

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			err = u.SetTransactionFailed(ctx, logger, req, senderTransaction)
			logger.Panic(p)
		}
		if err != nil {
			tx.Rollback()
			err = u.SetTransactionFailed(ctx, logger, req, senderTransaction)
			logger.Error(err)
		} else {
			tx.Commit()
		}
	}()

	recipientBalance, err := u.userBalanceRepository.FindByUserID(ctx, req.TargetUserID)
	if err != nil {
		logger.Error(err)
		return err
	}

	var (
		recipientBalanceBefore = recipientBalance.Balance
		recipientBalanceAfter  = recipientBalanceBefore.Add(senderTransaction.Amount)
	)

	recipientTransaction := &entity.UserTransaction{
		UserID:          req.UserID,
		TransactionType: constant.UserTransactionTypeCredit,
		ActionType:      constant.UserTransactionActionTypeTransfer,
		Amount:          senderTransaction.Amount,
		BalanceBefore:   recipientBalanceBefore,
		BalanceAfter:    recipientBalanceAfter,
		Remarks:         null.StringFromPtr(req.Remarks),
		Status:          constant.UserTransactionStatusSuccess,
	}
	err = u.userTransactionRepository.Upsert(ctx, recipientTransaction)
	if err != nil {
		logger.Error(err)
		return err
	}

	senderTransaction.Status = constant.UserTransactionStatusSuccess
	err = u.userTransactionRepository.Upsert(ctx, senderTransaction)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = u.userTransferHistoryRepository.Upsert(ctx, &entity.UserTransferHistory{
		RecipientTransactionID: recipientTransaction.ID,
		SenderTransactionID:    senderTransaction.ID,
		CreatedAt:              req.CreatedAt,
		UpdatedAt:              req.CreatedAt,
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (u *UseCase) SetTransactionFailed(ctx context.Context, logger *logrus.Entry, req request.ProcessTransferReq, senderTransaction *entity.UserTransaction) error {
	senderTransaction.Status = constant.UserTransactionStatusFailed
	err := u.userTransactionRepository.Upsert(ctx, senderTransaction)
	if err != nil {
		logger.Error(err)
		return err
	}

	userBalance, err := u.userBalanceRepository.FindByUserID(ctx, req.UserID)
	if err != nil {
		logger.Error(err)
		return err
	}

	userBalance.Balance = senderTransaction.BalanceBefore
	err = u.userBalanceRepository.Upsert(ctx, userBalance)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

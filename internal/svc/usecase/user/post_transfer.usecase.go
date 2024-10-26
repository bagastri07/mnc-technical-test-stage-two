package user

import (
	"context"
	"time"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/cachekey"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/task"
	"github.com/go-redsync/redsync/v4"
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) PostTransfer(ctx context.Context, req request.PostUserTransferReq) (*response.PostUserTransferResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
		tx                 = u.db.WithContext(ctx).Begin()
		userTransactionKey = cachekey.NewUserTransactionKey(req.UserID)
		mutex              = util.RedSyncObtainLock(u.redSync, userTransactionKey, redsync.WithExpiry(constant.LockUserTransferTimeout))
		err                error
	)

	util.RedSyncLock(ctx, mutex)
	defer util.RedSyncUnlock(ctx, mutex)

	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logger.Panic(p)
		}
		if err != nil {
			tx.Rollback()
			logger.Error(err)
		} else {
			tx.Commit()
		}
	}()

	userBalance, err := u.userBalanceRepository.FindByUserID(ctx, req.UserID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var (
		paymentAmountDec = decimal.NewFromFloat(req.Amount)
		balanceBefore    = userBalance.Balance
		balanceAfter     = balanceBefore.Sub(paymentAmountDec)
	)

	if balanceAfter.LessThan(decimal.Zero) {
		return nil, constant.ErrBalanceIsNotEnough
	}

	userBalance.Balance = balanceAfter
	err = u.userBalanceRepository.Upsert(ctx, userBalance)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	userTransaction := &entity.UserTransaction{
		UserID:          req.UserID,
		TransactionType: constant.UserTransactionTypeDebit,
		ActionType:      constant.UserTransactionActionTypeTransfer,
		Amount:          paymentAmountDec,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    balanceAfter,
		Remarks:         null.StringFromPtr(req.Remarks),
		Status:          constant.UserTransactionStatusPending,
	}
	err = u.userTransactionRepository.Upsert(ctx, userTransaction)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = util.ProcessPayloadAndEnqueueTask(ctx, u.asynqClient, task.AsynqProcessUserTransferTask, request.ProcessTransferReq{
		UserID:              req.UserID,
		TargetUserID:        req.TargetUserID,
		SenderTransactionID: userTransaction.ID,
		Remarks:             req.Remarks,
		CreatedAt:           time.Now(),
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.PostUserTransferResp{
		TransferID:    userTransaction.ID,
		Amount:        req.Amount,
		Remarks:       null.StringFromPtr(req.Remarks),
		BalanceBefore: balanceBefore.InexactFloat64(),
		BalanceAfter:  balanceAfter.InexactFloat64(),
		CreatedAt:     userTransaction.CreatedAt,
	}, nil
}

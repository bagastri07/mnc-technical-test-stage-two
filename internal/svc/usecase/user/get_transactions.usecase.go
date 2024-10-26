package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) GetTransactions(ctx context.Context, req request.GetUserTransactionsReq) ([]response.GetUserTransactionsResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	transactions, err := u.userTransactionRepository.GetByUserID(ctx, req.UserID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	items := make([]response.GetUserTransactionsResp, 0)
	for _, transaction := range transactions {
		primaryID := transaction.ID
		item := response.GetUserTransactionsResp{
			Status:          transaction.Status,
			UserID:          transaction.UserID,
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount.InexactFloat64(),
			Remarks:         transaction.Remarks,
			BalanceBefore:   transaction.BalanceBefore.InexactFloat64(),
			BalanceAfter:    transaction.BalanceAfter.InexactFloat64(),
			CreatedAt:       transaction.CreatedAt,
		}

		switch transaction.ActionType {
		case constant.UserTransactionActionTypePayment:
			item.PaymentID = &primaryID
		case constant.UserTransactionActionTypeTransfer:
			item.TransferID = &primaryID
		case constant.UserTransactionActionTypeTopUp:
			item.TopUpID = &primaryID
		}

		items = append(items, item)
	}

	return items, nil
}

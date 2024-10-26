package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func (h *Handler) ProcessTransfer(ctx context.Context, t *asynq.Task) error {
	var (
		logger = logrus.WithFields(logrus.Fields{})
		req    = request.ProcessTransferReq{}
	)

	ctx = util.SetDefaultLoggerValueToCtx(ctx, logger, constant.UserClaimKeys)

	err := util.BindingAsynqPayload(t, &req)
	if err != nil {
		logger.Error(err)
		return err
	}

	return h.userUseCase.ProcessTransfer(ctx, req)
}

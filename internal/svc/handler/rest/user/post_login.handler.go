package user

import (
	"net/http"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) PostLogin(c *gin.Context) {
	var (
		logger = logrus.WithFields(logrus.Fields{})
		ctx    = util.SetDefaultLoggerValueToCtx(c.Request.Context(), logger, constant.UserClaimKeys)
		req    = request.PostUserLoginReq{}
	)

	if err := c.ShouldBind(&req); err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	resp, err := h.userUseCase.PostLogin(ctx, req)
	if err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.CommonResultResp{
		Status: constant.SuccessMessage,
		Result: resp,
	})
}

package user

import (
	"net/http"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/claim"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetInfo(c *gin.Context) {
	var (
		logger = logrus.WithFields(logrus.Fields{})
		ctx    = util.SetDefaultLoggerValueToCtx(c.Request.Context(), logger, constant.UserClaimKeys)
		req    = request.GetUserInfoReq{}
	)

	userClaim, err := util.BindingFromContext[claim.User](ctx, constant.UserClaimKeys, nil)
	if err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	req.Username = userClaim.Username

	resp, err := h.userUseCase.GetInfo(ctx, req)
	if err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

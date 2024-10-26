package user

import (
	"net/http"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Register(c *gin.Context) {
	var (
		logger = logrus.WithFields(logrus.Fields{})
		ctx    = util.SetDefaultLoggerValueToCtx(c.Request.Context(), logger, constant.UserClaimKeys)
		req    = request.UserRegisterReq{}
	)

	if err := c.ShouldBind(&req); err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	resp, err := h.userUseCase.Register(ctx, req)
	if err != nil {
		logger.Error(err)
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

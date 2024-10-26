package healthcheck

import (
	"net/http"

	"github.com/bagastri07/gin-boilerplate-service/internal/config"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, response.MessageResp{
		Message: config.Env.App.Quote,
	})
}

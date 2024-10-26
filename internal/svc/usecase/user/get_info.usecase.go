package user

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/util"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/guregu/null/v5"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) GetInfo(ctx context.Context, req request.GetUserInfoReq) (*response.GetUserInfoResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.GetUserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		FullName: null.NewString(user.FullName.String, user.FullName.Valid),
	}, nil
}

package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) GetInfo(ctx context.Context, req request.GetUserInfoReq) (*response.GetUserInfoResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByPhoneNumber(ctx, req.Username)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.GetUserInfoResp{
		ID: user.ID,
		//Username: user.FirstName,
		//FullName: null.NewString(user.FullName.String, user.FullName.Valid),
	}, nil
}

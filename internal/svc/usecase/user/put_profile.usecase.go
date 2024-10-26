package user

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/guregu/null/v5"
	"github.com/sirupsen/logrus"
)

func (u *UseCase) PutProfile(ctx context.Context, req request.PutUserProfileReq) (*response.PutUserProfileResp, error) {
	var (
		logger = util.GetLoggerFromCtx(ctx).WithFields(logrus.Fields{
			"req": util.Dump(req),
		})
	)

	user, err := u.userRepository.FindByID(ctx, req.UserID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = null.StringFromPtr(req.Address)

	err = u.userRepository.Upsert(ctx, user)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &response.PutUserProfileResp{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

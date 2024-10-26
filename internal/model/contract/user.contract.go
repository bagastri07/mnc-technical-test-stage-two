package contract

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.User, error)
	FindByPhoneNumber(ctx context.Context, username string) (*entity.User, error)
	Upsert(ctx context.Context, user *entity.User) error
}

type UserUseCase interface {
	PostRegister(ctx context.Context, req request.PostUserRegisterReq) (*response.PostUserRegisterResp, error)
	PostLogin(ctx context.Context, req request.PostUserLoginReq) (*response.PostLoginUserResp, error)
	PostTopUp(ctx context.Context, req request.PostUserTopUpReq) (*response.PostUserTopUpResp, error)
	PostPayment(ctx context.Context, req request.PostUserPaymentReq) (*response.PostUserPaymentResp, error)
	PostTransfer(ctx context.Context, req request.PostUserTransferReq) (*response.PostUserTransferResp, error)
	ProcessTransfer(ctx context.Context, req request.ProcessTransferReq) error
	GetTransactions(ctx context.Context, req request.GetUserTransactionsReq) ([]response.GetUserTransactionsResp, error)
	PutProfile(ctx context.Context, req request.PutUserProfileReq) (*response.PutUserProfileResp, error)
}

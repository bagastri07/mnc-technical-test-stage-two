package contract

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/model/entity"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	Upsert(ctx context.Context, user *entity.User) error
}

type UserUseCase interface {
	Register(ctx context.Context, req request.UserRegisterReq) (*response.UserIDResp, error)
	Login(ctx context.Context, req request.UserLoginReq) (*response.TokenResponse, error)
	GetInfo(ctx context.Context, req request.GetUserInfoReq) (*response.GetUserInfoResp, error)
}

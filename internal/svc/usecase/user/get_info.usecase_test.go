package user

import (
	"context"
	"testing"

	mock_contract "github.com/bagastri07/gin-boilerplate-service/internal/model/contract/mock"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/entity"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/request"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/response"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"
)

func Test_useCase_GetInfo(t *testing.T) {
	var (
		ctrl               = gomock.NewController(t)
		mockUserRepository = mock_contract.NewMockUserRepository(ctrl)

		mockUUID = uuid.New()
	)
	type args struct {
		ctx context.Context
		req request.GetUserInfoReq
	}
	tests := []struct {
		name         string
		args         args
		expectedResp *response.GetUserInfoResp
		expectedErr  error
		prepareMock  func()
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				req: request.GetUserInfoReq{
					Username: "username",
				},
			},
			expectedErr: nil,
			expectedResp: &response.GetUserInfoResp{
				ID:       mockUUID,
				Username: "username",
			},
			prepareMock: func() {
				mockUserRepository.EXPECT().
					FindByUsername(gomock.Any(), "username").
					Times(1).
					Return(&entity.User{
						ID:       mockUUID,
						Username: "username",
					}, nil)
			},
		},
		{
			name: "find user by username err",
			args: args{
				ctx: context.Background(),
				req: request.GetUserInfoReq{
					Username: "username",
				},
			},
			expectedErr:  gorm.ErrInvalidDB,
			expectedResp: nil,
			prepareMock: func() {
				mockUserRepository.EXPECT().
					FindByUsername(gomock.Any(), "username").
					Times(1).
					Return(nil, gorm.ErrInvalidDB)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()

			uc := New().
				WithUserRepository(mockUserRepository)

			resp, err := uc.GetInfo(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

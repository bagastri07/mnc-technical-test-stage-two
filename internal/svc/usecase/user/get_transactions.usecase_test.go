package user

import (
	"context"
	"testing"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	mock_contract "github.com/bagastri07/mnc-technical-test-stage-two/internal/model/contract/mock"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"
)

func TestUseCase_GetTransactions(t *testing.T) {
	var (
		ctrl                          = gomock.NewController(t)
		mockUserTransactionRepository = mock_contract.NewMockUserTransactionRepository(ctrl)

		mockUUID = uuid.New()
	)

	type args struct {
		ctx context.Context
		req request.GetUserTransactionsReq
	}
	tests := []struct {
		name         string
		args         args
		expectedResp []response.GetUserTransactionsResp
		expectedErr  error
		prepareMock  func()
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				req: request.GetUserTransactionsReq{
					UserID: mockUUID,
				},
			},
			expectedResp: []response.GetUserTransactionsResp{
				{
					PaymentID: &mockUUID,
				},
				{
					TopUpID: &mockUUID,
				},
				{
					TransferID: &mockUUID,
				},
			},
			expectedErr: nil,
			prepareMock: func() {
				mockUserTransactionRepository.EXPECT().
					GetByUserID(gomock.Any(), mockUUID).
					Times(1).
					Return([]entity.UserTransaction{
						{
							ID:         mockUUID,
							ActionType: constant.UserTransactionActionTypePayment,
						},
						{
							ID:         mockUUID,
							ActionType: constant.UserTransactionActionTypeTopUp,
						},
						{
							ID:         mockUUID,
							ActionType: constant.UserTransactionActionTypeTransfer,
						},
					}, nil)
			},
		},
		{
			name: "get transaction err",
			args: args{
				ctx: context.Background(),
				req: request.GetUserTransactionsReq{
					UserID: mockUUID,
				},
			},
			expectedResp: nil,
			expectedErr:  gorm.ErrInvalidValue,
			prepareMock: func() {
				mockUserTransactionRepository.EXPECT().
					GetByUserID(gomock.Any(), mockUUID).
					Times(1).
					Return(nil, gorm.ErrInvalidValue)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareMock != nil {
				tt.prepareMock()
			}

			svc := New().
				WithUserTransactionRepository(mockUserTransactionRepository)

			resp, err := svc.GetTransactions(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

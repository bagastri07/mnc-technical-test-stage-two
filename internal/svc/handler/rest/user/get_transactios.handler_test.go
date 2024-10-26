package user

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/util"
	mock_contract "github.com/bagastri07/mnc-technical-test-stage-two/internal/model/contract/mock"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/customerror"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/request"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestHandler_GetTransactions(t *testing.T) {
	var (
		ctrl            = gomock.NewController(t)
		mockUserUsecase = mock_contract.NewMockUserUseCase(ctrl)

		mockUUID = uuid.New()

		expectedResp = response.CommonResultResp{
			Status: constant.SuccessMessage,
			Result: []response.GetUserTransactionsResp{
				{
					PaymentID: &mockUUID,
				},
			},
		}
	)

	tests := []struct {
		name             string
		expectedErr      *customerror.CustomError
		prepareMock      func()
		isMissingContext bool
	}{
		{
			name:        "normal",
			expectedErr: nil,
			prepareMock: func() {
				mockUserUsecase.EXPECT().
					GetTransactions(gomock.Any(), request.GetUserTransactionsReq{
						UserID: mockUUID,
					}).
					Times(1).
					Return([]response.GetUserTransactionsResp{
						{
							PaymentID: &mockUUID,
						},
					}, nil)
			},
		},
		{
			name:        "err",
			expectedErr: constant.ErrDataNotFound,
			prepareMock: func() {
				mockUserUsecase.EXPECT().
					GetTransactions(gomock.Any(), request.GetUserTransactionsReq{
						UserID: mockUUID,
					}).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			gin.SetMode(gin.TestMode)

			reqResult, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/v1/transactions", nil)
			if err != nil {
				t.Fatal(err)
			}

			reqResult.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = reqResult

			if !tt.isMissingContext {
				ctx := util.NewMockContext(c.Request.Context(), map[constant.CtxKey]any{
					constant.IDKey: mockUUID,
				})

				c.Request = c.Request.WithContext(ctx)
			}

			handler := New().
				WithUserUseCase(mockUserUsecase)

			handler.GetTransactions(c)

			b, _ := io.ReadAll(w.Body)

			if w.Code != http.StatusOK {
				var errResp customerror.CustomError
				err := json.Unmarshal(b, &errResp)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, *tt.expectedErr, errResp)
			} else {
				var resp response.CommonResultResp

				err := json.Unmarshal(b, &resp)
				if err != nil {
					t.Fatal(err)
				}

				result, err := bindToStructsUsingJSON(resp.Result)
				assert.NoError(t, err)

				assert.Equal(t, expectedResp.Result, result)
			}
		})
	}
}

func bindToStructsUsingJSON(data any) ([]response.GetUserTransactionsResp, error) {
	// Convert []interface{} to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to []MyStruct
	var result []response.GetUserTransactionsResp
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

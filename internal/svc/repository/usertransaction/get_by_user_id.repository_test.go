package usertransaction

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository_GetByUserID(t *testing.T) {
	var (
		db, mock, _ = sqlmock.New()
		mockUUID    = uuid.New()

		query = `SELECT * FROM "user_transactions" WHERE user_id = $1 ORDER BY created_at DESC`
	)

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	tests := []struct {
		name         string
		args         args
		expectedResp []entity.UserTransaction
		expectedErr  error
		prepareMock  func()
	}{
		{
			name: "normal",
			args: args{
				ctx:    context.Background(),
				userID: mockUUID,
			},
			expectedResp: []entity.UserTransaction{{ID: mockUUID}},
			expectedErr:  nil,
			prepareMock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(mockUUID)
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(mockUUID).
					WillReturnRows(rows)
			},
		},
		{
			name: "err",
			args: args{
				ctx:    context.Background(),
				userID: mockUUID,
			},
			expectedResp: nil,
			expectedErr:  gorm.ErrInvalidValue,
			prepareMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(mockUUID).
					WillReturnError(gorm.ErrInvalidValue)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareMock != nil {
				tt.prepareMock()
			}

			gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatal(err)
			}

			r := New().
				WithPostgresqlDB(gdb)

			resp, err := r.GetByUserID(tt.args.ctx, tt.args.userID)
			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

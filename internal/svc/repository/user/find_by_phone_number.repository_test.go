package user

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

func TestRepository_FindByPhoneNumber(t *testing.T) {
	var (
		db, mock, _ = sqlmock.New()
		mockUUID    = uuid.New()
		mockPhone   = "+62622"

		query = `SELECT * FROM "users" WHERE phone_number = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`
	)

	type args struct {
		ctx         context.Context
		phoneNumber string
	}

	tests := []struct {
		name         string
		args         args
		expectedResp *entity.User
		expectedErr  error
		prepareMock  func()
	}{
		{
			name: "normal",
			args: args{
				ctx:         context.Background(),
				phoneNumber: mockPhone,
			},
			expectedResp: &entity.User{ID: mockUUID},
			expectedErr:  nil,
			prepareMock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(mockUUID)
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(mockPhone, 1).
					WillReturnRows(rows)
			},
		},
		{
			name: "err",
			args: args{
				ctx:         context.Background(),
				phoneNumber: mockPhone,
			},
			expectedResp: nil,
			expectedErr:  gorm.ErrInvalidValue,
			prepareMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(mockPhone, 1).
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

			resp, err := r.FindByPhoneNumber(tt.args.ctx, tt.args.phoneNumber)
			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

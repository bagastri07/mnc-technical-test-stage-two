package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository_Upsert(t *testing.T) {
	var (
		db   *sql.DB
		mock sqlmock.Sqlmock

		mockUUID      = uuid.New()
		mockStr       = "test"
		userForUpdate = entity.User{
			ID: mockUUID,
		}

		userForInsert = entity.User{
			FirstName: mockStr,
		}

		queryInsert = `INSERT INTO "users" ("first_name","last_name","phone_number","pin","address","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`
		queryUpdate = `UPDATE "users" SET "first_name"=$1,"last_name"=$2,"phone_number"=$3,"pin"=$4,"address"=$5,"created_at"=$6,"updated_at"=$7,"deleted_at"=$8 WHERE "users"."deleted_at" IS NULL AND "id" = $9`
	)

	type args struct {
		ctx  context.Context
		user *entity.User
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
		prepareMock func()
	}{
		{
			name: "normal update",
			args: args{
				ctx:  context.Background(),
				user: &userForUpdate,
			},
			expectedErr: nil,
			prepareMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).
					WithArgs(
						userForUpdate.FirstName,
						userForUpdate.LastName,
						userForUpdate.PhoneNumber,
						userForUpdate.PIN,
						userForUpdate.Address,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						userForUpdate.ID,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "normal insert",
			args: args{
				ctx:  context.Background(),
				user: &userForInsert,
			},
			expectedErr: nil,
			prepareMock: func() {
				rows := sqlmock.NewRows([]string{"id", "vessel_id"}).AddRow(mockUUID, mockUUID)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(queryInsert)).
					WithArgs(
						userForInsert.FirstName,
						userForInsert.LastName,
						userForInsert.PhoneNumber,
						userForInsert.PIN,
						userForInsert.Address,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
					).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
		{
			name: "err",
			args: args{
				ctx:  context.Background(),
				user: &userForUpdate,
			},
			expectedErr: gorm.ErrDryRunModeUnsupported,
			prepareMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).
					WithArgs(
						userForUpdate.FirstName,
						userForUpdate.LastName,
						userForUpdate.PhoneNumber,
						userForUpdate.PIN,
						userForUpdate.Address,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						userForUpdate.ID,
					).
					WillReturnError(gorm.ErrDryRunModeUnsupported)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ = sqlmock.New()
			gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatal(err)
			}

			tt.prepareMock()

			r := New().
				WithPostgresqlDB(gdb)

			err = r.Upsert(tt.args.ctx, tt.args.user)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

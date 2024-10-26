package contract

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
)

type UserTransactionRepository interface {
	Upsert(ctx context.Context, userTransaction *entity.UserTransaction) error
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.UserTransaction, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserTransaction, error)
}

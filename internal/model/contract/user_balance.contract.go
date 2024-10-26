package contract

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	"github.com/google/uuid"
)

type UserBalanceRepository interface {
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserBalance, error)
	Upsert(ctx context.Context, userBalance *entity.UserBalance) error
}

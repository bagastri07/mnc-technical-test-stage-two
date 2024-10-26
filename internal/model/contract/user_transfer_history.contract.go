package contract

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
)

type UserTransferHistoryRepository interface {
	Upsert(ctx context.Context, userTransferHistory *entity.UserTransferHistory) error
}

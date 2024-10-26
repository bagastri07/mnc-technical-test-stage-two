package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserBalance struct {
	ID        uuid.UUID `gorm:"default:uuid_generate_v4()"`
	UserID    uuid.UUID
	Balance   decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}

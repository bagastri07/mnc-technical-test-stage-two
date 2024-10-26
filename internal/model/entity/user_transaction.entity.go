package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/shopspring/decimal"
)

type UserTransaction struct {
	ID              uuid.UUID `gorm:"default:uuid_generate_v4()"`
	UserID          uuid.UUID
	TransactionType string
	ActionType      string
	Amount          decimal.Decimal
	BalanceBefore   decimal.Decimal
	BalanceAfter    decimal.Decimal
	Remarks         null.String
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

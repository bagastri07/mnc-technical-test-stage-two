package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserTransferHistory struct {
	ID                     uuid.UUID `gorm:"default:uuid_generate_v4()"`
	RecipientTransactionID uuid.UUID
	SenderTransactionID    uuid.UUID
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

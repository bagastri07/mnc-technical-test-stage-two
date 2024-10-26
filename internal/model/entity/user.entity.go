package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"default:uuid_generate_v4()"`
	FirstName   string
	LastName    string
	PhoneNumber string
	PIN         string
	Address     null.String
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

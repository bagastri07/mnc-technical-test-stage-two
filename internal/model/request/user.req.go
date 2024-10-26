package request

import (
	"time"

	"github.com/google/uuid"
)

type PostUserLoginReq struct {
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	PIN         string `json:"pin" binding:"required"`
}

type PostUserRegisterReq struct {
	FirstName   string  `json:"first_name" binding:"required"`
	LastName    string  `json:"last_name" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,e164"`
	Address     *string `json:"address" binding:"omitempty,max=255"`
	PIN         string  `json:"pin" binding:"required,len=6"`
}

type GetUserInfoReq struct {
	Username string
}

type PostUserTopUpReq struct {
	UserID uuid.UUID `json:"-"`
	Amount float64   `json:"amount" binding:"required,min=1"`
}

type PostUserPaymentReq struct {
	UserID  uuid.UUID `json:"-"`
	Amount  float64   `json:"amount" binding:"required,min=1"`
	Remarks *string   `json:"remarks" binding:"omitempty,max=255"`
}

type PostUserTransferReq struct {
	UserID       uuid.UUID `json:"-"`
	TargetUserID uuid.UUID `json:"target_user_id" binding:"required"`
	Amount       float64   `json:"amount" binding:"required,min=1"`
	Remarks      *string   `json:"remarks" binding:"omitempty,max=255"`
}

type ProcessTransferReq struct {
	UserID              uuid.UUID `json:"user_id"`
	TargetUserID        uuid.UUID `json:"target_user_id"`
	SenderTransactionID uuid.UUID `json:"sender_transaction_id"`
	Remarks             *string   `json:"remarks"`
	CreatedAt           time.Time
}

type GetUserTransactionsReq struct {
	UserID uuid.UUID `json:"-"`
}

type PutUserProfileReq struct {
	UserID    uuid.UUID `json:"-"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Address   *string   `json:"address" binding:"omitempty,max=255"`
}

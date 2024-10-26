package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type UserIDResp struct {
	UserID uuid.UUID `json:"user_id"`
}

type PostUserRegisterResp struct {
	UserID      uuid.UUID   `json:"user_id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	PhoneNumber string      `json:"phone_number"`
	Address     null.String `json:"address"`
	CreatedAt   time.Time   `json:"created_at"`
}

type GetUserInfoResp struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	FullName null.String `json:"full_name"`
}

type PostUserTopUpResp struct {
	TopUpID       uuid.UUID `json:"top_up_id"`
	AmountTopUp   float64   `json:"amount_top_up"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedAt     time.Time `json:"created_at"`
}

type PostUserPaymentResp struct {
	PaymentID     uuid.UUID   `json:"payment_id"`
	Amount        float64     `json:"amount"`
	Remarks       null.String `json:"remarks"`
	BalanceBefore float64     `json:"balance_before"`
	BalanceAfter  float64     `json:"balance_after"`
	CreatedAt     time.Time   `json:"created_at"`
}

type PostUserTransferResp struct {
	TransferID    uuid.UUID   `json:"transfer_id"`
	Amount        float64     `json:"amount"`
	Remarks       null.String `json:"remarks"`
	BalanceBefore float64     `json:"balance_before"`
	BalanceAfter  float64     `json:"balance_after"`
	CreatedAt     time.Time   `json:"created_at"`
}

type GetUserTransactionsResp struct {
	TransferID      *uuid.UUID  `json:"transfer_id,omitempty"`
	PaymentID       *uuid.UUID  `json:"payment_id,omitempty"`
	TopUpID         *uuid.UUID  `json:"top_up_id,omitempty"`
	Status          string      `json:"status"`
	UserID          uuid.UUID   `json:"user_id"`
	TransactionType string      `json:"transaction_type"`
	Amount          float64     `json:"amount"`
	Remarks         null.String `json:"remarks"`
	BalanceBefore   float64     `json:"balance_before"`
	BalanceAfter    float64     `json:"balance_after"`
	CreatedAt       time.Time   `json:"created_at"`
}

type PutUserProfileResp struct {
	UserID    uuid.UUID   `json:"user_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Address   null.String `json:"address"`
	UpdatedAt time.Time   `json:"updated_at"`
}

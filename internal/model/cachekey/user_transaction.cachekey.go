package cachekey

import (
	"fmt"

	"github.com/google/uuid"
)

func NewUserTransactionKey(userID uuid.UUID) string {
	return fmt.Sprintf("user_transaction:%s", userID.String())
}

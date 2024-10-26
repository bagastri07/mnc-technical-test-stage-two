package response

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type UserIDResp struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetUserInfoResp struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	FullName null.String `json:"full_name"`
}

package util

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"gorm.io/gorm"
)

func NewTxContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, constant.DB, tx)
}

func GetTxFromContext(ctx context.Context, defaultTx *gorm.DB) *gorm.DB {
	txVal := ctx.Value(constant.DB)
	tx, ok := txVal.(*gorm.DB)
	if !ok {
		return defaultTx
	}
	return tx
}

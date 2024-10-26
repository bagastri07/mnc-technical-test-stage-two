package util

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
)

func NewMockContext(ctx context.Context, dataMap map[constant.CtxKey]any) context.Context {
	for key, value := range dataMap {
		ctx = context.WithValue(ctx, key, value) //nolint
	}

	return ctx
}

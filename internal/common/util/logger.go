package util

import (
	"context"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/sirupsen/logrus"
)

func SetDefaultLoggerValueToCtx(ctx context.Context, logger *logrus.Entry, keys []constant.CtxKey) context.Context {
	for _, key := range keys {
		val := ctx.Value(key)
		if val != nil {
			logger = logger.WithField(string(key), val)
		}
	}

	return context.WithValue(ctx, constant.LoggerKey, logger)
}

func GetLoggerFromCtx(ctx context.Context) *logrus.Entry {
	val := ctx.Value(constant.LoggerKey)

	if logger, ok := val.(*logrus.Entry); ok {
		return logger
	}

	return logrus.WithContext(ctx)
}

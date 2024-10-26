package util

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	"github.com/sirupsen/logrus"
)

func RedSyncObtainLock(redsyncClient *redsync.Redsync, key string, options ...redsync.Option) *redsync.Mutex {
	mutex := redsyncClient.NewMutex(key, options...)
	return mutex
}

func RedSyncLock(ctx context.Context, mutex *redsync.Mutex) {
	if err := mutex.LockContext(ctx); err != nil {
		logrus.Panic("locking failed", err)
	}
}

func RedSyncUnlock(ctx context.Context, mutex *redsync.Mutex) {
	if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
		logrus.Panic("unlock failed", err)
	}
}

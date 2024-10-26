package distributed

import (
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/task"
	"github.com/bagastri07/mnc-technical-test-stage-two/internal/svc/handler/distributed/user"
	"github.com/hibiken/asynq"
)

type ServeMuxBuilder struct {
	mux *asynq.ServeMux
}

func NewServeMuxBuilder() *ServeMuxBuilder {
	return &ServeMuxBuilder{
		mux: asynq.NewServeMux(),
	}
}

func (b *ServeMuxBuilder) WithUserHandler(userHandler *user.Handler) *ServeMuxBuilder {
	b.mux.HandleFunc(task.AsynqProcessUserTransferTask, userHandler.ProcessTransfer)
	return b
}

func (b *ServeMuxBuilder) Build() *asynq.ServeMux {
	return b.mux
}

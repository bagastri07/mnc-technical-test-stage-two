package util

import (
	"context"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/common/constant"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func BindingAsynqPayload(t *asynq.Task, out any) error {
	err := json.Unmarshal(t.Payload(), &out)
	if err != nil {
		return err
	}
	return nil
}

func processPayloadToAsynqTask(task string, body any) (*asynq.Task, error) {
	byteData, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(task, byteData), nil
}

func enqueueTask(ctx context.Context, asynqClient *asynq.Client, task *asynq.Task, opts ...asynq.Option) error {
	result, err := asynqClient.EnqueueContext(ctx, task, opts...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"task":    task.Type(),
			"payload": string(task.Payload()),
		}).Error(err)
		return err
	}

	logrus.WithFields(logrus.Fields{
		"taskID":  result.ID,
		"task:":   task.Type(),
		"payload": string(task.Payload()),
	}).Info("Enqueue Task Success ")

	return nil
}

func ProcessPayloadAndEnqueueTask(ctx context.Context, asynqClient *asynq.Client, task string, body any, opts ...asynq.Option) error {
	asynqTask, err := processPayloadToAsynqTask(task, body)
	if err != nil {
		return err
	}

	if len(opts) <= 0 {
		opts = []asynq.Option{
			asynq.Retention(constant.AsynqDefaultRetention),
			asynq.MaxRetry(constant.AsynqDefaultMaxRetry),
		}
	}

	return enqueueTask(ctx, asynqClient, asynqTask, opts...)
}

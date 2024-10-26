package distributed

import (
	"context"
	"time"

	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() asynq.MiddlewareFunc {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			var (
				start = time.Now()
			)

			logrus.Infof("Start processing: %q", t.Type())
			beatifyJSONPayload(t.Payload())
			err := h.ProcessTask(ctx, t)
			if err != nil {
				return err
			}
			logrus.Infof("Finished processing: %q taskID: %s Elapsed Time: %v\n", t.Type(), t.ResultWriter().TaskID(), time.Since(start))
			return nil
		})
	}
}

func beatifyJSONPayload(payload []byte) {
	var (
		payloadMap map[string]interface{}
	)

	err := json.Unmarshal(payload, &payloadMap)
	if err != nil {
		logrus.Infof("Task Payload:  %q", string(payload))
	} else {
		prettyPayload, _ := json.MarshalIndent(json.RawMessage(payload), "", "    ")
		logrus.Infof("Task Payload:  %+v", string(prettyPayload))
	}
}

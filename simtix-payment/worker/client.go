package worker

import (
	"github.com/hibiken/asynq"
)

type workerClient struct {
	client *asynq.Client
}

func (w *workerClient) Close() error {
	return w.client.Close()
}

func (w *workerClient) Enqueue(t *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return w.client.Enqueue(t, opts...)
}

func NewWorkerClient(redisAddr string) workerClient {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	return workerClient{
		client: client,
	}
}

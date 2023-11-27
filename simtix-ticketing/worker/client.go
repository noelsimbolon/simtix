package worker

import (
	"github.com/hibiken/asynq"
)

//var Module = fx.Module("workerClient", fx.Options(fx.Provide(NewWorkerClient)))

type WorkerClient struct {
	client *asynq.Client
}

func (w *WorkerClient) Close() error {
	return w.client.Close()
}

func (w *WorkerClient) Enqueue(t *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return w.client.Enqueue(t, opts...)
}

// to do: take from config
func NewWorkerClient() *WorkerClient {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	return &WorkerClient{
		client: client,
	}
}

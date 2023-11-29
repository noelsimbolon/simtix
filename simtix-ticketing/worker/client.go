package worker

import (
	"fmt"
	"github.com/hibiken/asynq"
	"simtix-ticketing/config"
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
func NewWorkerClient(config *config.Config) *WorkerClient {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort)})
	return &WorkerClient{
		client: client,
	}
}

package worker

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"simtix/lib"
	"simtix/utils/logger"
	"simtix/worker/handlers"
	"simtix/worker/tasks"
)

type WorkerServer struct {
	srv            *asynq.Server
	router         *asynq.ServeMux
	paymentHandler *handlers.MakePaymentHandler
}

func NewServer(config *lib.Config, paymentHandler *handlers.MakePaymentHandler) *WorkerServer {
	var server WorkerServer
	redisConnOpt := asynq.RedisClientOpt{Addr: config.RedisAddress}
	server.srv = asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(server.HandleError),
		},
	)

	server.router = asynq.NewServeMux()
	server.router.HandleFunc(tasks.TypeMakePaymentTask, paymentHandler.HandleMakePaymentTask())

	return &server
}

func (s *WorkerServer) Run() {
	if err := s.srv.Run(s.router); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (s *WorkerServer) HandleError(ctx context.Context, task *asynq.Task, err error) {
	retried, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)
	logger.Log.Info(fmt.Sprintf("retries:%d, maxretry: %d", retried, maxRetry))
	if retried >= maxRetry {
		err = fmt.Errorf("retry exhausted for task %s: %w", task.Type(), err)
		if task.Type() == tasks.TypeMakePaymentTask {
			err = s.paymentHandler.HandleError(task)
		}
		logger.Log.Error(err)
	}
}

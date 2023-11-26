package worker

import (
	"github.com/hibiken/asynq"
	"log"
	"ticketing/worker/handlers"
	"ticketing/worker/tasks"
)

type WorkerServer struct {
	srv        *asynq.Server
	router     *asynq.ServeMux
	pdfHandler *handlers.GeneratePdfHandler
}

func NewServer(pdfHandler *handlers.GeneratePdfHandler) *WorkerServer {
	var server WorkerServer
	redisConnOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	server.srv = asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			//ErrorHandler: asynq.ErrorHandlerFunc(server.HandleError),
		},
	)

	server.router = asynq.NewServeMux()
	server.router.HandleFunc(tasks.TypeGeneratePdfTask, pdfHandler.HandleGeneratePdf())

	return &server
}

func (s *WorkerServer) Run() {
	if err := s.srv.Run(s.router); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

//func (s *WorkerServer) HandleError(ctx context.Context, task *asynq.Task, err error) {
//	retried, _ := asynq.GetRetryCount(ctx)
//	maxRetry, _ := asynq.GetMaxRetry(ctx)
//	if retried >= maxRetry {
//		//err = fmt.Errorf("retry exhausted for task %s: %w", task.Type(), err)
//		if task.Type() == tasks.TypeMakePaymentTask {
//			err = s.paymentHandler.HandleError(task)
//		}
//
//	}
//}

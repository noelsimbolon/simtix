package worker

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"simtix-ticketing/config"
	"simtix-ticketing/worker/handlers"
	"simtix-ticketing/worker/tasks"
)

type WorkerServer struct {
	srv        *asynq.Server
	router     *asynq.ServeMux
	pdfHandler *handlers.GeneratePdfHandler
}

func NewServer(pdfHandler *handlers.GeneratePdfHandler, config *config.Config) *WorkerServer {
	var server WorkerServer
	redisConnOpt := asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort)}
	server.srv = asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
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

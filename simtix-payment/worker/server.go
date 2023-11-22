package worker

import (
	"github.com/hibiken/asynq"
	"log"
)

type WorkerServer struct {
	srv    *asynq.Server
	router *asynq.ServeMux
}

func NewServer(redisAddr string, queues map[string]int) WorkerServer {
	const asynqConcurrency int = 5
	redisConnOpt := asynq.RedisClientOpt{Addr: redisAddr}
	srv := asynq.NewServer(
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

	router := asynq.NewServeMux()
	// to implement, paymentSimulationHandler is injected
	//router.HandleFunc(tasks.PaymentSimulation, paymentSimulationHandler.Execute())
	return WorkerServer{
		srv:    srv,
		router: router,
	}
}

func (s *WorkerServer) Run() {
	if err := s.srv.Run(s.router); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"simtix-ticketing/config"
	"simtix-ticketing/route"
	"syscall"
	"time"
)

var Module = fx.Module("server", fx.Provide(NewServer))

type Server struct {
	engine *gin.Engine
	port   int
	routes  *route.Routes
}

func NewServer(config *config.Config, routes *route.Routes) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	routes.Setup(engine.Group("/api"))

	return &Server{
		engine: engine,
		port: config.ServicePort,
		routes: routes,
	}
}

func (s *Server) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(
			signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT,
		)
		<-signals
		cancel()
	}()

	address := fmt.Sprintf(":%d", s.port)

	server := &http.Server{
		Addr:    address,
		Handler: s.engine,
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.Go(
		func() error {
			return s.startServer(server)
		},
	)

	group.Go(
		func() error {
			<-groupCtx.Done()
			ctx, cancelTo := context.WithTimeout(context.Background(), time.Second*30)
			defer cancelTo()
			fmt.Println("Shutting down server")
			return server.Shutdown(ctx)
		},
	)

	if err := group.Wait(); err != nil {
		fmt.Println("Failed to shutdown server")
	}
}

func (s *Server) startServer(server *http.Server) error {
	fmt.Println(fmt.Sprintf("Server started on port %d", s.port))
	return server.ListenAndServe()
}

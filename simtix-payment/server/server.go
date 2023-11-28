package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"simtix/lib"
	"simtix/middlewares"
	"simtix/routes"
	"simtix/utils/logger"
	"syscall"
	"time"
)

var Module = fx.Module("server", fx.Provide(NewServer))

type Server struct {
	engine *gin.Engine
	port   int
	routes *routes.Routes
}

func NewServer(routes *routes.Routes, config *lib.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.LoggingMiddleware())
	routes.Setup(engine.Group("/api/payment"))
	return &Server{
		engine: engine,
		port:   config.ServicePort,
		routes: routes,
	}
}

func (s *Server) SetupRoutes() {
	s.engine.GET(
		"/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello, Gin server!"})
		},
	)
}

func (s *Server) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file")
	}

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
			logger.Log.Info("Shutting down server")
			return server.Shutdown(ctx)
		},
	)

	if err := group.Wait(); err != nil {
		logger.Log.Error("Failed to shutdown server")
	}
}

func (s *Server) startServer(server *http.Server) error {
	logger.Log.Info(fmt.Sprintf("Server started on port %d", s.port))
	return server.ListenAndServe()
}

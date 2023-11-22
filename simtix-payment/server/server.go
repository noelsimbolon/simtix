package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simtix/lib"
	"simtix/routes"
	"strconv"
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
	engine := gin.Default()
	routes.Setup(engine.Group("/api"))
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
		log.Println("Error loading .env file")
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(
			signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT,
		)
		<-signals
		cancel()
	}()

	port := getPort()

	address := fmt.Sprintf(":%d", port)

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
			log.Print(`Shutting down payment service`)
			return server.Shutdown(ctx)
		},
	)

	if err := group.Wait(); err != nil {
		log.Printf(`Exiting due to %v`, err)
	}
}

func (s *Server) startServer(server *http.Server) error {
	log.Printf("Starting server on port %d", s.port)
	return server.ListenAndServe()
}

func getPort() int {
	val, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080
	}
	port, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Error converting PORT to integer: %v", err)
		return 8080
	}
	return port
}

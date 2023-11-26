package lib

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
	"simtix/utils/logger"
	"time"
)

type Config struct {
	ServiceHost        string        `envconfig:"SERVICE_HOST" required:"true"`
	ServiceState       int           `envconfig:"SERVICE_STATE" required:"true" default:"0"`
	ServiceName        string        `envconfig:"SERVICE_NAME" required:"true"`
	ServicePort        int           `envconfig:"SERVICE_PORT" default:"8000" required:"true"`
	GracefullyDuration time.Duration `envconfig:"GRACEFULLY_DURATION" default:"5s"`
	RedisAddress       string        `envconfig:"REDIS_ADDRESS"`
	TicketingEndpoint  string        `envconfig:"TICKETING_ENDPOINT"`

	LogLevel         string `envconfig:"LOG_LEVEL" default:"INFO"`
	DatabaseHost     string `envconfig:"DB_HOST" required:"true"`
	DatabasePort     string `envconfig:"DB_PORT" required:"true"`
	DatabaseUsername string `envconfig:"DB_USERNAME" required:"true"`
	DatabasePassword string `envconfig:"DB_PASSWORD" required:"true"`
	DatabaseName     string `envconfig:"DB_NAME" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	logger.Log.Info(fmt.Sprintf("Loading env from file: %s", filename))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
		return &config, nil
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}

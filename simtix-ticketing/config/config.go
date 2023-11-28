package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Module("config", fx.Options(fx.Provide(NewConfig)))

type Config struct {
	ServiceHost     string `envconfig:"SERVICE_HOST" required:"true"`
	ServiceState    int    `envconfig:"SERVICE_STATE" required:"true" default:"0"`
	ServiceName     string `envconfig:"SERVICE_NAME" required:"true"`
	ServicePort     int    `envconfig:"SERVICE_PORT" default:"8000" required:"true"`
	PaymentEndpoint string `envconfig:"PAYMENT_ENDPOINT" required:"true"`

	RedisHost string `envconfig:"REDIS_HOST"`
	RedisPort string `envconfig:"REDIS_PORT"`

	DatabaseHost     string `envconfig:"DB_HOST" required:"true"`
	DatabasePort     int    `envconfig:"DB_PORT" required:"true"`
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

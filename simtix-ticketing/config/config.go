package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

var Module = fx.Module("config", fx.Options(fx.Provide(NewConfig)))

type Config struct {
	ServiceHost  string `envconfig:"SERVICE_HOST" required:"true"`
	ServiceState int    `envconfig:"SERVICE_STATE" required:"true" default:"0"`
	ServiceName  string `envconfig:"SERVICE_NAME" required:"true"`
	ServicePort  int    `envconfig:"SERVICE_PORT" default:"8000" required:"true"`

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

	err := envconfig.Process("", &config)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to read environment variables.")
	}

	return &config, nil
}

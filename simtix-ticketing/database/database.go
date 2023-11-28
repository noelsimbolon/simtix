package database

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simtix-ticketing/config"
)

var Module = fx.Module("database", fx.Options(fx.Provide(NewDatabase)))

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	host := config.DatabaseHost
	port := config.DatabasePort
	dbname := config.DatabaseName
	user := config.DatabaseUsername
	password := config.DatabasePassword

	dsn := fmt.Sprintf("host=%s port=%d  dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, dbname, user, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to establish connection to database.")
	}

	return &Database{
		DB: db,
	}, nil
}

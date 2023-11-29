package amqp

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"simtix-ticketing/config"
)

type AmqpClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewAmqpClient(config *config.Config) (*AmqpClient, error) {
	log.Print(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%d", config.AmqpUser, config.AmqpPassword,
			config.AmqpHost, config.AmqpPort,
		),
	)
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%d", config.AmqpUser, config.AmqpPassword,
			config.AmqpHost, config.AmqpPort,
		),
	)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &AmqpClient{
		conn:    conn,
		channel: ch,
	}, nil
}

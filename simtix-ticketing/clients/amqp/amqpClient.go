package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewAmqpClient() (*AmqpClient, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
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

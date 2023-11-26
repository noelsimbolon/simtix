package amqp

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type BookingProcessedPayload struct {
	Pattern string      `json:"pattern""`
	Data    interface{} `json:"data"`
}

func (c *AmqpClient) declareBookingProcessedQueue() (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		"CLIENT_QUEUE",
		true,
		false,
		false,
		false,
		nil,
	)
}

// TO DO:
// pass booking on the param
func (c *AmqpClient) SendBookingProcessedMessage() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	queue, err := c.declareBookingProcessedQueue()
	if err != nil {
		return err
	}
	body := BookingProcessedPayload{
		Pattern: "BOOKING_PROCESSED",
		Data:    "tes tes aja duluh",
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = c.channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			//Body: json.Marshal(booking),
		},
	)
	log.Print(err)
	return err
}

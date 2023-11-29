package amqp

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"simtix-ticketing/model"
	"time"
)

type BookingDataPayload struct {
	BookingID  string           `json:"bookingID"`
	PdfUrl     string           `json:"pdfUrl"`
	SeatStatus model.SeatStatus `json:"seatStatus"`
}

type BookingProcessedPayload struct {
	Pattern string             `json:"pattern"`
	Data    BookingDataPayload `json:"data"`
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

func (c *AmqpClient) SendBookingProcessedMessage(data BookingDataPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	queue, err := c.declareBookingProcessedQueue()
	if err != nil {
		return err
	}
	body := BookingProcessedPayload{
		Pattern: "BOOKING_PROCESSED",
		Data:    data,
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

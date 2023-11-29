package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
	"log"
	"simtix-ticketing/clients/amqp"
	"simtix-ticketing/worker/tasks"
	"time"
)

type GeneratePdfHandler struct {
	amqpClient *amqp.AmqpClient
}

func NewGeneratePdfHandler(amqpClient *amqp.AmqpClient) *GeneratePdfHandler {
	return &GeneratePdfHandler{
		amqpClient: amqpClient,
	}
}

// to do pass the booking object here
func (h *GeneratePdfHandler) HandleGeneratePdf() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		log.Print("HANDLEEEE PDFFFFFF")
		payload, err := h.unmarshalPayload(t)

		if err != nil {
			log.Print("DISIIINIII ERORRRRRR")
			return err
		}

		pdf := gofpdf.New("P", "mm", "A4", "")

		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Ticket Booking Successful")
		pdf.Ln(10)

		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(40, 10, "BOOKING DETAILS")
		pdf.Ln(10)

		h.addTicketDetails(pdf, payload)

		err = h.generateQrCode(pdf)
		if err != nil {
			log.Print("DISIIINIII ERORRRRRR 2")
			return err
		}

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		pdfPath := fmt.Sprintf("public/tickets/BOOKING_%s_%s.pdf", timestamp, payload.BookingID)
		err = pdf.OutputFileAndClose(pdfPath)
		if err != nil {
			log.Print("DISIIINIII ERORRRRRR 3")
			return err
		}

		bookingProcessedData := amqp.BookingDataPayload{
			BookingID:  payload.BookingID,
			PdfUrl:     pdfPath,
			SeatStatus: payload.Seat.Status,
		}
		err = h.amqpClient.SendBookingProcessedMessage(bookingProcessedData)
		if err != nil {
			return err
		}

		return nil
	}
}

func (h *GeneratePdfHandler) addTicketDetails(pdf *gofpdf.Fpdf, data *tasks.GeneratePdfPayload) {
	seat := data.Seat

	pdf.SetFont("Arial", "", 14)
	col1X := 10.0
	col2X := 70.0
	lineHeight := 10.0
	addDetail := func(label, value string) {
		pdf.SetX(col1X)
		pdf.Cell(0, lineHeight, label)
		pdf.SetX(col2X)
		pdf.Cell(0, lineHeight, value)
		pdf.Ln(-1)
	}

	addDetail("Event:", seat.Event.EventName)
	addDetail("Date:", seat.Event.EventTime.Format(time.RFC822))
	addDetail("Seat:", fmt.Sprintf("Row %s, Number %d", seat.SeatRow, seat.SeatNumber))
	addDetail("Price:", seat.Price.String())
	addDetail("Booking ID:", data.BookingID)
}

func (h *GeneratePdfHandler) generateQrCode(pdf *gofpdf.Fpdf) error {
	// encode booking id
	qrCode, err := qrcode.Encode("example text here", qrcode.Medium, 256)
	if err != nil {
		return err
	}

	pageWidth, pageHeight := pdf.GetPageSize()
	qrCodeSize := 50.0
	qrCodeX := (pageWidth - qrCodeSize) / 2
	qrCodeY := (pageHeight - qrCodeSize) / 2
	qrCodeReader := bytes.NewReader(qrCode)

	pdf.RegisterImageOptionsReader("QRCodeImage", gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader)
	pdf.ImageOptions("QRCodeImage", qrCodeX, qrCodeY, qrCodeSize, qrCodeSize, false, gofpdf.ImageOptions{}, 0, "")

	return nil
}

func (h *GeneratePdfHandler) unmarshalPayload(t *asynq.Task) (*tasks.GeneratePdfPayload, error) {
	var payload tasks.GeneratePdfPayload

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Print("Failed to unmarshal payload")
		return nil, err
	}
	return &payload, nil
}

package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
	"time"
)

type GeneratePdfHandler struct {
}

func NewGeneratePdfHandler() *GeneratePdfHandler {
	return &GeneratePdfHandler{}
}

// to do pass the booking object here
func (h *GeneratePdfHandler) HandleGeneratePdf() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {

		pdf := gofpdf.New("P", "mm", "A4", "")

		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Ticket Booking Successful")
		pdf.Ln(10)

		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(40, 10, "BOOKING DETAILS")

		h.addTicketDetails(pdf)

		err := h.generateQrCode(pdf)
		if err != nil {
			return err
		}

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		pdfPath := fmt.Sprintf("public/tickets/BOOKING_%s_%s.pdf", timestamp, "AKSNXFJL")
		err = pdf.OutputFileAndClose(pdfPath)
		if err != nil {
			return err
		}

		return nil
	}
}

func (h *GeneratePdfHandler) addTicketDetails(pdf *gofpdf.Fpdf) {
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

	addDetail("Event:", "Concert")
	addDetail("Date:", "January 1, 2024")
	addDetail("Seat:", "Section A, Row 1, Seat 1")
	addDetail("Price:", "$50.00")
	addDetail("Booking ID:", "123456")
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

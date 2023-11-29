package main

import (
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
	"time"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Ticket Booking Successful")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "BOOKING DETAILS")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 14)
	col1X := 10.0
	col2X := 70.0
	lineHeight := 10.0

	// Add ticket details (dummy text) in two columns
	addDetail := func(label, value string) {
		pdf.SetX(col1X)
		pdf.Cell(0, lineHeight, label)
		pdf.SetX(col2X)
		pdf.Cell(0, lineHeight, value)
		pdf.Ln(-1)
	}

	// Add ticket details
	addDetail("Event:", "Concert")
	addDetail("Date:", "January 1, 2024")
	addDetail("Seat:", "Section A, Row 1, Seat 1")
	addDetail("Price:", "$50.00")
	addDetail("Booking ID:", "123456")

	err := generateQrCode(pdf)
	if err != nil {
		log.Fatal(err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	pdfPath := fmt.Sprintf("static/tickets/BOOKING_%s_%s.pdf", timestamp, "AKSNXFJL")
	err = pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", pdfPath)
}

// saveQrCodeImage saves the QR code image to the specified file path.
func saveQrCodeImage(filePath string, qrCode []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(qrCode)
	if err != nil {
		return err
	}

	return nil
}

func generateQrCode(pdf *gofpdf.Fpdf) error {
	// encode booking id
	qrCode, err := qrcode.Encode("example text here", qrcode.Medium, 256)
	if err != nil {
		return err
	}

	pageWidth, _ := pdf.GetPageSize()
	qrCodeSize := 50.0
	qrCodeX := (pageWidth - qrCodeSize) / 2
	qrCodeReader := bytes.NewReader(qrCode)

	pdf.RegisterImageOptionsReader("QRCodeImage", gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader)
	pdf.ImageOptions("QRCodeImage", qrCodeX, pdf.GetY(), qrCodeSize, qrCodeSize, false, gofpdf.ImageOptions{}, 0, "")

	return nil
}

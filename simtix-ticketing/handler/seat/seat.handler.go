package seat

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simtix-ticketing/config"
	"simtix-ticketing/model/dto"
	"simtix-ticketing/service/seat"
)

type SeatHandler interface {
	GetSeatsByEvent(c *gin.Context)
	GetSeatByID(c *gin.Context)
	PostSeat(c *gin.Context)
	SeatWebhook(c *gin.Context)
	PatchSeatForBooking(c *gin.Context)
}

type SeatHandlerImpl struct {
	service       seat.SeatService
	webhookSecret string
}

func NewSeatHandler(service seat.SeatService, config *config.Config) *SeatHandlerImpl {
	return &SeatHandlerImpl{
		service:       service,
		webhookSecret: config.WebhookSecret,
	}
}

func (h *SeatHandlerImpl) GetSeatsByEvent(c *gin.Context) {
	eventID := c.Query("eventID")
	if eventID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Event ID parameter is required"})
		return
	}
	seats, err := h.service.GetSeatsByEventID(eventID)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, seats)
}

func (h *SeatHandlerImpl) GetSeatByID(c *gin.Context) {
	seatID := c.Param("seatID")
	if seatID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Seat ID parameter is required"})
		return
	}
	seat, err := h.service.GetSeatByID(seatID)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
	return
}

func (h *SeatHandlerImpl) PostSeat(c *gin.Context) {
	var dto dto.CreateSeatDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, custErr := h.service.CreateSeat(dto)
	if custErr != nil {
		c.AbortWithStatusJSON(custErr.StatusCode, gin.H{"error": custErr.Err.Error()})
		return
	}

	c.JSON(http.StatusCreated, seat)
	return
}

func (h *SeatHandlerImpl) PatchSeatForBooking(c *gin.Context) {
	var dto dto.BookSeatDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, custErr := h.service.BookSeat(dto)
	if custErr != nil {
		c.AbortWithStatusJSON(custErr.StatusCode, gin.H{"error": custErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
	return
}

func (h *SeatHandlerImpl) SeatWebhook(c *gin.Context) {
	signature := c.GetHeader("X-Webhook-Signature")
	log.Print(signature)
	body, err := c.GetRawData()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	if err := h.checkWebhookSignature(signature, body); err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid webhook signature"})
		return
	}

	var payload dto.UpdateSeatStatusDto
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(payload)
	log.Print(h.webhookSecret)
	_, custErr := h.service.UpdateSeatStatus(payload)
	if custErr != nil {
		c.JSON(custErr.StatusCode, gin.H{"error": custErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed successfully"})
}

func (h *SeatHandlerImpl) checkWebhookSignature(signature string, body []byte) error {
	calculatedSignature, err := h.calculateHMAC(body)
	if err != nil {
		return fmt.Errorf("failed to calculate HMAC: %v", err)
	}

	if signature != calculatedSignature {
		return fmt.Errorf("invalid webhook signature")
	}
	return nil
}

func (h *SeatHandlerImpl) calculateHMAC(body []byte) (string, error) {
	key := []byte(h.webhookSecret)
	hashed := hmac.New(sha256.New, key)
	hashed.Write(body)
	signature := base64.StdEncoding.EncodeToString(hashed.Sum(nil))
	return signature, nil
}

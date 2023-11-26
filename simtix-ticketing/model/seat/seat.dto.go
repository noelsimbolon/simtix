package seat

type CreateSeatDto struct {
	EventID string     `json:"eventID" binding:"required"`
	Status  SeatStatus `json:"status" binding:"required"`
}

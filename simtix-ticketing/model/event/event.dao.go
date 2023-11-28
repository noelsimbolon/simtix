package event

type GetAllEventsDao struct {
	Events []Event `json:"events"`
}

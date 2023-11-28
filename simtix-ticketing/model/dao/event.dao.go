package dao

import "simtix-ticketing/model"

type GetAllEventsDao struct {
	Events []model.Event `json:"events"`
}

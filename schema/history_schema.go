package schema

import "gin-framework-services/models"

type HistorySchema struct {
	Message string `json:message`
	Status int `json:status`
	Data []models.History `json:data`
}

type HistoryByID struct {
	Message string `json:message`
	Status int `json:status`
	Data models.HistoryByID `json:data`
}

type HistoryCRUD struct {
	Message string `json:message`
	Status int `json:status`
	Data models.History `json:data`
}
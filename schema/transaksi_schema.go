package schema

import "gin-framework-services/models"

type TransaksiSchema struct {
	Message string `json:message`
	Status int `json:status`
	Data   []models.Transaksi `json:data`
}

type TransaksiByID struct {
	Message string `json:message`
	Status int `json:status`
	Data   models.TransaksiByID `json:data`
}

type TransaksiCRUD struct {
	Message string `json:message`
	Status int `json:status`
	Data   models.Transaksi `json:data`
}


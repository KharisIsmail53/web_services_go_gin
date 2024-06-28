package schema

import (
	"gin-framework-services/models"
)

type ZakatSchema struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    []models.Zakat `json:"data"`
}

type ZakatCRUD struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    models.Zakat `json:"data"`
}

type ZakatByID struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    models.Zakat `json:"data"`
}
package schema

import (
	"gin-framework-services/models"
)

type BerasSchema struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    []models.Beras `json:"data"`
}

type BerasCRUD struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    models.Beras `json:"data"`
}

type BerasByID struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    models.Beras `json:"data"`
}


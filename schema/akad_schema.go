package schema

import (
	"gin-framework-services/models"
)

type AkadSchema struct {
	Message string       `json:"message"`
	Status  int          `json:"status"`
	Data    []models.Akad `json:"data"`
}

type AkadCRUD struct {
	AkadSchema
	Data models.Akad `json:"data"`
}

type AkadByID struct {
	AkadSchema
	Data models.Akad `json:"data"`
}
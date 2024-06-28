package schema

import "gin-framework-services/models"

type KategoriSchema struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    []models.Kategori `json:"data"`
}

type KategoriCRUD struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    models.Kategori `json:"data"`
}

type KategoriByID struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    models.Kategori `json:"data"`
}
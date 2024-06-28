package schema

import "gin-framework-services/models"

type BarangSchema struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    []models.Barang `json:"data"`
}

type BarangCRUD struct{
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    models.Barang `json:"data"`
}

type BarangByID struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Data    models.BarangByID `json:"data"`
}


package models

import (
	"github.com/google/uuid"
)

type History struct {
	BaseUUIDModel
	IDTransaksi uuid.UUID `gorm:"type:uuid;index" json:"id_transaksi"`
	Tahun       int       `gorm:"type:int;not null" json:"tahun"`
}

type HistoryByID struct {
	History
	TransaksiByID
}

func (History) TableName() string {
	return "history"
}
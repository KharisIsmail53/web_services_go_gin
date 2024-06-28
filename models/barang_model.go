package models

import (
	"github.com/google/uuid"
)

type Barang struct {
    BaseUUIDModel
    IDKategori uuid.UUID `gorm:"type:uuid;index" json:"id_kategori,omitempty"`
    NamaBarang string    `gorm:"type:varchar(255);not null" json:"nama_barang"`
    Stock      int       `gorm:"not null" json:"stock"`
}

type BarangByID struct{
    Barang
    Kategori Kategori `gorm:"foreignKey:IDKategori" json:"kategori"`
}

func (Barang) TableName() string {
    return "barang"
}

func (BarangByID) TableName() string {
    return "barang"
}
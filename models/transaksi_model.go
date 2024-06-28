package models

import (
	"github.com/google/uuid"
)

type Transaksi struct {
	BaseUUIDModel
	IDBeras        *uuid.UUID `gorm:"type:uuid;index" json:"id_beras"`
	IDAkad         *uuid.UUID `gorm:"type:uuid;index" json:"id_akad"`
	IDZakat        *uuid.UUID `gorm:"type:uuid;index" json:"id_zakat"`
	JumlahLiteran   *float32  `gorm:"type:float;not null" json:"jumlah_literan"`
	JumlahKeluarga  *int       `gorm:"type:int;not null" json:"jumlah_keluarga"`
	JumlahUang      *int       `gorm:"type:int;not null" json:"jumlah_uang"`
}

type TransaksiByID struct {
	Transaksi
	Beras *Beras `gorm:"foreignKey:IDBeras" json:"beras"`
	Akad  *Akad  `gorm:"foreignKey:IDAkad" json:"akad"`
	Zakat *Zakat `gorm:"foreignKey:IDZakat" json:"zakat"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}

func (TransaksiByID) TableName() string {
	return "transaksi"
}
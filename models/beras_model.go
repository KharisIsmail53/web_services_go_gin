package models

type Beras struct {
	BaseUUIDModel
	Deskripsi string `json:"deskripsi" gorm:"type:varchar(255);not null"`
	HargaBeras int `json:"harga_beras" gorm:"type:int;not null"`
	Stock float32 `json:"stock" gorm:"type:float;not null"`
}

func (Beras) TableName() string {
	return "beras"
}
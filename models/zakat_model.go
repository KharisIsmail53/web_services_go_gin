package models

type Zakat struct {
	BaseUUIDModel
	JenisZakat string `json:"jenis_zakat" gorm:"type:varchar(255);not null"`
}

func (Zakat) TableName() string {
	return "zakat"
}
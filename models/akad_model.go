package models

type Akad struct {
	BaseUUIDModel
	JenisAkad string `json:"jenis_akad" gorm:"type:varchar(255);not null;"`
}

func (Akad) TableName() string {
	return "akad"
}
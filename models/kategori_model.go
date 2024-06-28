package models

type Kategori struct {
	BaseUUIDModel
	NamaKategori string `json:"nama_kategori"`
	Satuan       string `json:"satuan"`
}

func (Kategori) TableName() string {
	return "kategori"
}

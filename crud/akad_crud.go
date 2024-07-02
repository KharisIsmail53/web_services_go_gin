package crud

import (
	"github.com/google/uuid"
	"gin-framework-services/models"
	"gorm.io/gorm"
)

func GetAkad(db *gorm.DB, id uuid.UUID) (models.Akad, error) {
	var akad models.Akad
	if err := db.Where("id = ?", id).First(&akad).Error; err != nil {
		return models.Akad{}, err
	}
	return akad, nil
}
package crud

import (
	"github.com/google/uuid"
	"gin-framework-services/models"
	"gorm.io/gorm"
)

func GetBeras(db *gorm.DB, id uuid.UUID) (models.Beras, error) {
	var beras models.Beras
	if err := db.Where("id = ?", id).First(&beras).Error; err != nil {
		return models.Beras{}, err
	}
	return beras, nil
}
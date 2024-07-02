package crud

import (
	"github.com/google/uuid"
	"gin-framework-services/models"
	"gorm.io/gorm"
)

func GetZakat(db *gorm.DB, id uuid.UUID) (models.Zakat, error) {
	var zakat models.Zakat
	if err := db.Where("id = ?", id).First(&zakat).Error; err != nil {
		return models.Zakat{}, err
	}
	return zakat, nil
}
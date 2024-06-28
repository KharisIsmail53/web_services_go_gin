package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseUUIDModel struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedByID *uuid.UUID `gorm:"type:uuid;index;default:'c903429b-9e8c-4ee7-85fc-12d3a9e0a2fb'" json:"created_by_id"`
	UpdatedByID *uuid.UUID `gorm:"type:uuid;index;default:'c903429b-9e8c-4ee7-85fc-12d3a9e0a2fb'" json:"updated_by_id"`
}

func (base *BaseUUIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return
}

func (base *BaseUUIDModel) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now()
	return
}

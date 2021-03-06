package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Uuid uuid.UUID `gorm:"type:char(36)" `
}

// BeforeCreate These functions are called before creating Base
func (base *Base) BeforeCreate(db *gorm.DB) error {

	base.Uuid = uuid.Must(uuid.NewV4())

	return nil
}

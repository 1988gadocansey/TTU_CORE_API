package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Uuid uuid.UUID `gorm:"primaryKey" gorm:"type:char(36)"`
	//ID   uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
}

// BeforeCreate These functions are called before creating Base
func (base *Base) BeforeCreate(db *gorm.DB) error {
	id := uuid.Must(uuid.NewV4())
	db.Model(base).Update("uuid", id)

	return nil
}

package models

import "gorm.io/gorm"

type Disability struct {
	Base
	gorm.Model

	Name      string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	StudentID uint8
	Student   Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
}

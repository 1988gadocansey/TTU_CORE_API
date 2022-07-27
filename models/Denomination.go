package models

import "gorm.io/gorm"

type Denomination struct {
	Base
	gorm.Model

	Name       string   `json:"Name" gorm:"uniqueIndex" validate:"required"`
	ReligionID uint8    ` json:"ReligionId"`
	Religion   Religion ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Religion"`
}

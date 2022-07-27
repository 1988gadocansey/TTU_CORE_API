package models

import "gorm.io/gorm"

type District struct {
	Base
	gorm.Model

	Name     string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	RegionID uint8  ` json:"RegionID"`
	Region   Region ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Region"`
}

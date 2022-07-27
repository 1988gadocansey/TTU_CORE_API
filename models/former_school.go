package models

import (
	"gorm.io/gorm"
)

type FormerSchool struct {
	Base
	gorm.Model
	Name       string `db:"Name" json:"Name" validate:"required"`
	DistrictID uint8
	District   District `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"District"`
}

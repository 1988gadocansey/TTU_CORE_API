package models

import "gorm.io/gorm"

type Examination struct {
	Base
	gorm.Model
	Name          string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	AffiliationID uint
	Body          Affiliation `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Body"`
}

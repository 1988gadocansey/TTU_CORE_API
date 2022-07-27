package models

import "gorm.io/gorm"

type Affiliation struct {
	Base
	gorm.Model

	Name   string `json:"Name" gorm:"uniqueIndex" validate:"required,lte=255"`
	Active bool   `json:"Active"  validate:"required"`

	//Programme []Programme `gorm:"ForeignKey:AffiliationID" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Affiliation"`
}

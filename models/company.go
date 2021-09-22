package models

import "gorm.io/gorm"

type CompanyCategory struct {
	Base
	gorm.Model

	Name    string    ` json:"Name" gorm:"uniqueIndex"`
	Company []Company `gorm:"ForeignKey:CompanyCategoryID" json:"Website,Companies" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}
type Company struct {
	Base
	gorm.Model

	Name              string `  json:"Name" gorm:"uniqueIndex" validate:"required,lte=255"`
	Description       string ` json:"Description,omitempty"`
	CompanyCategoryID int
	CompanyCategory   CompanyCategory `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"CompanyCategoryID"`
	Website           string          `gorm:"column:Website" json:"Website,omitempty"`
}

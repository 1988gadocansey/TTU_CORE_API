package models

import "gorm.io/gorm"

type Course struct {
	Base
	gorm.Model
	Name     string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	Type     string `json:"Type"  validate:"required"` // elective or core or coreAlt
	Code     string `json:"ShortCode" gorm:"uniqueIndex"  validate:"required"`
	Credit   uint8  `json:"Credit"`
	Semester uint8  `json:"Semester"`
}

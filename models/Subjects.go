package models

import "gorm.io/gorm"

type Subject struct {
	Base
	gorm.Model
	Name          string      `json:"Name" gorm:"uniqueIndex" validate:"required"`
	Type          string      `json:"Type"  validate:"required"` // elective or core or coreAlt
	ShortCode     string      `json:"ShortCode" gorm:"uniqueIndex"  validate:"required"`
	ExaminationID uint8       ` json:"ReligionId"`
	Examination   Examination ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Examination"`
}

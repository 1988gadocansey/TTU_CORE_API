package models

import (
	"gorm.io/gorm"
)

type Calender struct {
	Base
	gorm.Model

	Year     string `db:"Year" json:"Year" gorm:"uniqueIndex" validate:"required"`
	Semester string `db:"Semester" json:"Semester" gorm:"uniqueIndex" validate:"required"`
}

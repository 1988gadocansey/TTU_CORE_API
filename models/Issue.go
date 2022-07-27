package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	Base
	gorm.Model
	Name       string `db:"Name" json:"Name" validate:"required"`
	StudentID  uint
	Student    Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
	CalenderID uint
	Calender   Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
}

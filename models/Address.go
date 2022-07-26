package models

import (
	"gorm.io/gorm"
)

type Address struct {
	Base
	gorm.Model
	City      string `db:"City" json:"City" validate:"required"`
	Box       string `db:"Box" json:"Box"`
	Street    string `db:"Street" json:"Street" validate:"required"`
	GPost     string `db:"GPost" json:"GPost" validate:"required"`
	StudentID uint
	Student   Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
}

package models

import (
	"gorm.io/gorm"
)

type Address struct {
	Base
	gorm.Model

	Addresses string `db:"Addresses" json:"Addresses" `
	StudentID uint
	Student   Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
}

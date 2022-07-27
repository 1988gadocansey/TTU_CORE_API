package models

import "gorm.io/gorm"

type Clearance struct {
	Base
	gorm.Model
	Unit        string `json:"Office"  validate:"required"`
	Reason      string `json:"Reason"  validate:"required"`
	StudentID   uint8
	Student     Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Student"`
	CalenderID  uint8
	Calender    Calender `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
	ClearedByID uint8
	ClearedBy   User `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"ClearedBy"`
}

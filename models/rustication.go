package models

import (
	"gorm.io/gorm"
)

type Rustication struct {
	Base
	gorm.Model
	StudentID  uint8    `  json:"student_id,omitempty"`
	Student    Student  `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"student""`
	Duration   string   `json:"duration,omitempty"`
	Action     string   `json:"action,omitempty"`
	Reason     string   `json:"reason,omitempty"`
	Status     uint8    `json:"status"`
	CalenderID uint8    `json:"calender_id,omitempty"`
	Calender   Calender `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender""`
}

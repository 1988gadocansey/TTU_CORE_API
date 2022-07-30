package models

import (
	"gorm.io/gorm"
)

type RoomAllocation struct {
	Base
	gorm.Model
	StudentID  uint8    `json:"student_id,omitempty"`
	Student    Student  `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"student""`
	RoomID     uint8    `json:"room_id,omitempty"`
	Room       Room     `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"room""`
	Active     bool     `json:"active"`
	ISProtocol bool     `json:"is_protocol"`
	CalenderID uint8    `json:"calender_id,omitempty"`
	Calender   Calender `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender""`
}

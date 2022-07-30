package models

import (
	"gorm.io/gorm"
)

type Room struct {
	Base
	gorm.Model
	Gender         GenderData `json:"gender"`
	Floor          string     `json:"floor"` // first second third
	Bed            string     `json:"bed"`
	BedOrientation string     `json:"bed_orientation"` // top or down
	Capacity       uint8      `json:"capacity"`
	Reserved       bool       `json:"reserved"` //eg for src people...
	Status         uint8      `json:"status"`
	HallID         uint8      `json:"hall_id,omitempty"`
	Hall           Hall       `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"hall""`
}

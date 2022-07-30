package models

import (
	"gorm.io/gorm"
)

type MountedCourse struct {
	Base
	gorm.Model
	CalenderID  uint8
	Calender    Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender"`
	CourseID    uint8
	Course      Course ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"course"`
	ProgrammeID uint8
	Programme   Programme ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"programme"`
	Sync        bool      `json:"Sync"`
	Section     Sections  `json:"Sections"`
	LecturerID  uint8
	Lecturer    User ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"lecturer"`
}

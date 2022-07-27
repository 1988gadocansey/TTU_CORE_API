package models

import "gorm.io/gorm"

type AcademicRecord struct {
	Base
	gorm.Model
	MountedCourseID uint
	MountedCourse   MountedCourse ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"MountedCourse"`
	StudentID       uint
	Student         Student ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Student"`
	GradingID       uint
	Grading         GradingSystem ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Grading"`
	GradeStatusID   uint
	GradeStatus     GradeStatuses `json:"GradeStatus"`
	Resit           bool          `json:"Resit"`
	ResitStatus     ResitStatuses `json:"ResitStatus"`
	Sync            bool          `json:"Sync"`
	CalenderID      uint8
	Calender        Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
}

package models

import "gorm.io/gorm"

type AcademicRecord struct {
	Base
	gorm.Model
	MountedCourseID uint8
	MountedCourse   MountedCourse ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"MountedCourse"`
	StudentID       uint8
	Student         Student ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Student"`
	GradingID       uint8
	Grading         GradingSystem ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Grading"`
	GradeStatus     GradeStatuses `json:"GradeStatus"`
	Resit           bool          `json:"Resit"`
	ResitStatus     ResitStatuses `json:"ResitStatus"`
	Sync            bool          `json:"Sync"`
	CalenderID      uint8
	Calender        Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
}

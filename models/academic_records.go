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
	QuizOne         float32       `json:"QuizOne"  validate:"required"`
	QuizTwo         float32       `json:"QuizTwo"  validate:"required"`
	QuizThree       float32       `json:"QuizThree"  validate:"required"`
	MidSemOne       float32       `json:"MidSemOne"  validate:"required"`
	MidSemTwo       float32       `json:"MidSemTwo"`
	Project         float32       `json:"Project"`
	Total           float32       `json:"Total"`
	Exams           float32       `json:"Exams"  validate:"required"`
	GPoint          float32       `json:"GPoint"  validate:"required"`
	Grading         GradingSystem ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Grading"`
	GradeStatus     GradeStatuses `json:"GradeStatus"`
	YearGroup       string        `json:"YearGroup"`
	Resit           bool          `json:"Resit"`
	ResitStatus     ResitStatuses `json:"ResitStatus"`
	Sync            bool          `json:"Sync"`
	CalenderID      uint8
	Calender        Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
}

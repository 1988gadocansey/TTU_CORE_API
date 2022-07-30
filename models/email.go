package models

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	Email     string  `json:"email" gorm:"unique" validate:"required, email" json:"email,omitempty"`
	StudentID uint    `json:"student_id,omitempty"`
	Student   Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"student" bson:"student"`
}

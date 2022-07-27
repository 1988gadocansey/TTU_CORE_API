package models

import "gorm.io/gorm"

type StudentConstraint struct {
	Base
	gorm.Model

	TotalCredit        uint8   `json:"TotalCredit"  validate:"required,lte=25"`
	CreditDone         uint8   `json:"CreditDone"  validate:"required,lte=25"`
	CreditLeft         uint8   `json:"CreditLeft"  validate:"required,lte=25"`
	BiodataUpdated     bool    `json:"BiodataUpdated" validate:"required"`
	QualityAssurance   bool    `json:"QualityAssurance" validate:"required"`
	Liaison            bool    `json:"Liaison" validate:"required"`
	Registered         bool    `json:"Registered" validate:"required"`
	AllowedRegister    bool    `json:"AllowedRegister" validate:"required"`
	AllowedToSeeResult bool    `json:"AllowedToSeeResult" validate:"required"`
	StudentID          uint8   ` json:"StudentID"`
	Student            Student ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Student"`

	//SomeFlag bool    `gorm:"column:some_flag;not null;default:true"`

}

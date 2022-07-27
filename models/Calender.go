package models

import (
	"gorm.io/gorm"
	"time"
)

type Calender struct {
	Base
	gorm.Model
	Year                         string    `db:"Year" json:"Year" gorm:"uniqueIndex" validate:"required"`
	Semester                     uint8     `db:"Semester" json:"Semester"  validate:"required"`
	Running                      bool      `db:"Running" json:"Running" validate:"required"`
	OpenAdmission                bool      `db:"OpenAdmission" json:"OpenAdmission"  validate:"required"`
	OpenRegistration             bool      `db:"OpenRegistration" json:"OpenRegistration"  validate:"required"`
	OpenResult                   bool      `db:"OpenResult" json:"OpenResult"  validate:"required"`
	OpenQa                       bool      `db:"OpenQa" json:"OpenQa"  validate:"required"`
	OpenLa                       bool      `db:"OpenLa" json:"OpenLa"  validate:"required"`
	EnterMarks                   bool      `db:"EnterMarks" json:"EnterMarks"  validate:"required"`
	AdmissionFrom                time.Time `db:"AdmissionFrom" json:"AdmissionFrom"  validate:"required"`
	AdmissionTo                  time.Time `db:"AdmissionTo" json:"AdmissionTo"  validate:"required"`
	RegistrationFrom             time.Time `db:"RegistrationFrom" json:"RegistrationFrom"  validate:"required"`
	EnterMarksFrom               time.Time `db:"EnterMarksFrom" json:"EnterMarksFrom"  validate:"required"`
	FeePaymentFrom               time.Time `db:"FeePaymentFrom" json:"FeePaymentFrom"  validate:"required"`
	FeePaymentTo                 time.Time `db:"FeePaymentTo" json:"FeePaymentTo"  validate:"required"`
	SpecialRegistrationProgramme string    `db:"SpecialRegistrationProgramme" json:"SpecialRegistrationProgramme"`
	SpecialRegistrationLevel     string    `db:"SpecialRegistrationLevel" json:"SpecialRegistrationLevel"`
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type LiaisonData struct {
	Base
	gorm.Model
	StudentID            uint8   `  json:"student_id,omitempty"`
	Student              Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"student""`
	CompanyName          string  `  json:"company_name,omitempty"`
	CompanyAddress       string  `  json:"company_address,omitempty"`
	CompanyEmail         string  `  json:"company_email,omitempty"`
	CompanyPhone         string  `  json:"company_phone,omitempty"`
	CompanyTown          string  `  json:"company_town,omitempty"`
	CompanyExactLocation string  `  json:"company_exact_location,omitempty"`
	CompanyCategoryID    uint8
	CompanyCategory      CompanyCategory `json:"company_category,omitempty"`
	CompanyZoneID        uint8
	Zone                 Zone `json:"zone"`
	AddressToID          uint8
	AddressTo            Address   `json:"address_to,omitempty"`
	SupervisorID         uint8     `  json:"supervisor_id,omitempty"`
	Supervisor           User      `json:"supervisor" `
	Assumed              bool      `json:"assumed"`
	AssumptionDate       time.Time ` json:"assumption_date"`
	CalenderID           uint8     `json:"calender_id,omitempty"`
	Calender             Calender  `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender""`
}

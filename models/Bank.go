package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	Base
	gorm.Model
	Name          string `db:"Name" json:"Name" validate:"required"`
	AccountNumber string `db:"AccountNumber" json:"AccountNumber" validate:"required"`
	Active        bool   `db:"Active" json:"Active" validate:"required"`
}

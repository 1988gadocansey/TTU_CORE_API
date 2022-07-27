package models

import "gorm.io/gorm"

type Religion struct {
	Base
	gorm.Model

	Name string `json:"Name" gorm:"uniqueIndex" validate:"required"`
}

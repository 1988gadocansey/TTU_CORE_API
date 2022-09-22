package dto

import "gorm.io/gorm"

type StudentDTO struct {
	gorm.Model

	Name string `json:"Name" gorm:"uniqueIndex" validate:"required,lte=255"`
}

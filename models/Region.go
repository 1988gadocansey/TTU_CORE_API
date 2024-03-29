package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model

	Name string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	Slug string `json:"Slug" gorm:"uniqueIndex"`
}

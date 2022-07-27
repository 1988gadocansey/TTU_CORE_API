package models

import "gorm.io/gorm"

type Nationality struct {
	Base
	gorm.Model

	Name    string `json:"Name" gorm:"uniqueIndex" validate:"required"`
	Slug    string `json:"Slug" gorm:"uniqueIndex"`
	Country string `json:"Country" gorm:"uniqueIndex"`
}

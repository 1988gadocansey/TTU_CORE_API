package models

import "gorm.io/gorm"

type Hall struct {
	Base
	gorm.Model
	Name     string     `json:"Name" gorm:"uniqueIndex" validate:"required"`
	Slug     string     `json:"Slug" gorm:"uniqueIndex"`
	Type     GenderData ` json:"Type"`
	Capacity int32      ` json:"Capacity"`
}

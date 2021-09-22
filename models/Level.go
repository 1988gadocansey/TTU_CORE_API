package models

import "gorm.io/gorm"

type Level struct {
	Base
	gorm.Model

	Name string `json:"Name" gorm:"uniqueIndex" validate:"required,lte=255"`
	Slug string `json:"Slug" gorm:"uniqueIndex" validate:"required,lte=255"`
	//SomeFlag bool    `gorm:"column:some_flag;not null;default:true"`

	Student []Student `gorm:"ForeignKey:LevelID" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
}

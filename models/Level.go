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

func (level *Level) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Level{}).Count(&total)

	return total
}

func (level *Level) Take(db *gorm.DB, limit int, offset int) interface{} {
	var levels []Level

	db.Offset(offset).Limit(limit).Find(&level)

	return levels
}

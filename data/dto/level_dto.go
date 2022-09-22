package dto

import "gorm.io/gorm"

type LevelDTO struct {
	gorm.Model

	Name string `json:"Name" gorm:"uniqueIndex" validate:"required,lte=255"`
	Slug string `json:"Slug" gorm:"uniqueIndex" validate:"required,lte=255"`
	BaseDto
}

func (level *LevelDTO) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&LevelDTO{}).Count(&total)

	return total
}

func (level *LevelDTO) Take(db *gorm.DB, limit int, offset int) interface{} {
	var levels []LevelDTO

	db.Offset(offset).Limit(limit).Find(&levels)

	return levels
}

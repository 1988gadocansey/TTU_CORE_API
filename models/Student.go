package models

import (
	"gorm.io/gorm"
)

type Student struct {
	Base
	gorm.Model
	Stno        string `db:"Stno" json:"Stno" gorm:"uniqueIndex" validate:"required,Stno"`
	Indexno     string `db:"Indexno" json:"Indexno" gorm:"uniqueIndex" validate:"required,Indexno"`
	CustomId    string `db:"CustomId" json:"CustomId" gorm:"uniqueIndex"`
	Titles      AppTitle
	FirstName   string     `db:"FirstName" json:"FirstName"  validate:"required,FirstName"`
	LastName    string     `db:"LastName" json:"LastName"  validate:"required,LastName"`
	OtherName   string     `db:"OtherName" json:"OtherName"`
	Gender      GenderData `db:"Gender" json:"Gender"  validate:"required"`
	Email       string     `db:"Email" json:"Email" validate:"required,Email"`
	LevelID     uint8      ` json:"LevelID" validate:"required,LastName"`
	Level       Level      ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Levels"`
	ProgrammeID uint8      ` json:"ProgrammeID"`
	Programme   Programme  ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Programme"`
	Address     []Address  `gorm:"ForeignKey:StudentID" json:"Address" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}

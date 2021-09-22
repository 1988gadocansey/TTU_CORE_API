package models

import "gorm.io/gorm"

type Student struct {
	Base
	gorm.Model

	Stno        string `db:"Stno" json:"Stno" gorm:"uniqueIndex" validate:"required,Stno"`
	Indexno     string `db:"Indexno" json:"Indexno" gorm:"uniqueIndex" validate:"required,Indexno"`
	CustomId    string `db:"CustomId" json:"CustomId" gorm:"uniqueIndex"`
	FirstName   string `db:"FirstName" json:"FirstName"  validate:"required,FirstName"`
	LastName    string `db:"LastName" json:"LastName"  validate:"required,LastName"`
	OtherName   string `db:"OtherName" json:"OtherName"`
	LevelID     int    ` json:"LevelID"`
	Level       Level  ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Levels"`
	ProgrammeID int
	Programme   Programme ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Programme"`
	Address     []Address `gorm:"ForeignKey:StudentID" json:"Address" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}

//string `gorm:"unique;not null"`

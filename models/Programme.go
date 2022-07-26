package models

import "gorm.io/gorm"

type Faculty struct {
	Base
	gorm.Model
	Name string `db:"Name" json:"Name" validate:"required,Name"`
	Code string `db:"Code" json:"Code" validate:"required,Code"`
	// faculty has many departments
	Department []Department `gorm:"ForeignKey:FacultyID" json:"Departments" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}
type Department struct {
	Base
	gorm.Model
	Name      string  ` db:"Name" json:"Name" gorm:"uniqueIndex" validate:"required"`
	Code      string  `db:"Code" json:"Code" gorm:"uniqueIndex" validate:"required"`
	FacultyID uint    ` json:"FacultyID"`
	Faculty   Faculty ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Faculty"`

	// department has many programme
	Programme []Programme `gorm:"ForeignKey:DepartmentId" json:"Programme" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}

type Programme struct {
	Base
	gorm.Model
	Name         string     `db:"Name" json:"Name" gorm:"uniqueIndex" validate:"required,Name"`
	Slug         string     `db:"Slug" json:"Slug" gorm:"uniqueIndex" validate:"required,Slug"`
	DepartmentId uint       ` json:"DepartmentId"`
	Department   Department ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Department"`

	Duration string ` db:"Duration" json:"Duration"  validate:"required,Duration"`

	Student []Student `gorm:"ForeignKey:ProgrammeID" json:"Student" gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}

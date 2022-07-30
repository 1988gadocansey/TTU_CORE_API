package models

import (
	"gorm.io/gorm"
)

type Bill struct {
	Base
	gorm.Model
	LevelID        uint8
	Level          Level     ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Level"`
	Programme      Programme ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Programme"`
	CalenderID     uint8
	Calender       Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
	Services       float32  `db:"Services" json:"Services"`
	AUF            float32  `db:"AUF" json:"AUF"`
	DepartmentDues float32  `db:"DepartmentDues" json:"DepartmentDues"`
	Thesis         float32  `db:"Thesis" json:"Thesis"`
	Practicals     float32  `db:"Practical" json:"Practicals"`
	Compensation   float32  `db:"Compensation" json:"Compensation"`
	Src            float32  `db:"Src" json:"Src"`
	Liaison        float32  `db:"Liaison" json:"Liaison"`
	Cloth          float32  `db:"Cloth" json:"Cloth"`
	PPE            float32  `db:"PPE" json:"PPE"`
	Section        Sections `json:"section"`        // sandwich, regular, distance ,evening
	Type           float32  `db:"Type" json:"Type"` // ghanaian, foreign students

	Active bool `db:"Active" json:"Active" validate:"required"`
}

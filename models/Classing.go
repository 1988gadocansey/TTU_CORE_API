package models

import (
	"gorm.io/gorm"
)

type Classing struct {
	Base
	gorm.Model
	Lower                 float32 `db:"Lower" json:"Lower" validate:"required"`
	Upper                 float32 `db:"Upper" json:"Upper"`
	Class                 string  `db:"Class" json:"Class" validate:"required"`
	AwardingInstitutionID uint
	AwardingInstitution   Affiliation ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Institution"`
}

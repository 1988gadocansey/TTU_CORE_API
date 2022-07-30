package models

import (
	"gorm.io/gorm"
)

type Addresse struct {
	Base
	gorm.Model
	Name string `db:"Name" json:"Name" validate:"required"`
}

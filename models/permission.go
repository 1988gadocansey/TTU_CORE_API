package models

import "gorm.io/gorm"

type Permission struct {
	Base
	gorm.Model
	Name string `json:"name"`
}

package models

import "gorm.io/gorm"

type Zone struct {
	Base
	gorm.Model
	DistrictID uint8    `json:"district_id"`
	District   District `json:"district"`
	SubZone    string   `json:"sub_zone"  gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;"`
}

package models

import "gorm.io/gorm"

type GradingSystem struct {
	Base
	gorm.Model

	Grade         rune        `json:"Name"  validate:"required"`
	Lower         float32     `json:"Lower"  validate:"required"`
	Upper         float32     `json:"Upper"  validate:"required"`
	Value         float32     `json:"Value"  validate:"required"`
	Type          StudTypes   `json:"Type"  validate:"required"`
	AffiliationID uint8       ` json:"AffiliationId"`
	Affiliation   Affiliation ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Affiliation"`
}

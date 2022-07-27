package models

import (
	"gorm.io/gorm"
	"time"
)

type PaymentProduct struct {
	Base
	gorm.Model
	Name         string    `db:"Name" json:"Name" validate:"required"`
	Status       bool      `db:"Status" json:"Status" validate:"required"`
	Amount       float32   `db:"Amount" json:"Amount"`
	Banks        string    `db:"Banks" json:"Banks" validate:"required"`
	Code         string    `db:"Code" json:"Code"`
	PartPayment  bool      `db:"PartPayment" json:"PartPayment" validate:"required"`
	StartDate    time.Time `db:"StartDate" json:"StartDate" validate:"required"`
	EndDate      time.Time `db:"EndDate" json:"EndDate" validate:"required"`
	Instructions string    `db:"Instructions" json:"Instructions" validate:"required"`
}

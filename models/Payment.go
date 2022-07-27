package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	Base
	gorm.Model
	StudentID        uint8
	Student          Student `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Students"`
	LevelID          uint8
	Level            Level   ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Levels"`
	Amount           float32 `db:"Amount" json:"Amount"  validate:"required, numeric"`
	BankID           uint8
	Bank             Bank ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Bank"`
	PaymentTypeID    uint8
	PaymentType      PaymentProduct ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"PaymentType"`
	TransactionId    string         `db:"TransactionId" json:"TransactionId" `
	TransactionDate  time.Time      `db:"TransactionDate" json:"TransactionDate" `
	Source           PaymentSources `db:"Source" json:"Source"  validate:"required"`
	ReceiptNo        string         `db:"ReceiptNo" json:"ReceiptNo" `
	PaymentNarration string         `db:"PaymentNarration" json:"PaymentNarration" `
	Verified         bool           `db:"Verified" json:"Verified" `
	VerifiedByID     uint8
	VerifiedBy       User `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"VerifiedBy"`
	CalenderID       uint8
	Calender         Calender ` gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"Calender"`
}

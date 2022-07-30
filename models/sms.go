package models

import (
	"gorm.io/gorm"
	"time"
)

type Sms struct {
	Base
	gorm.Model

	Phone       string    `json:"phone"  validate:"required"`
	Message     string    `json:"message"`
	Status      bool      `json:"status"`
	Type        string    `json:"type"`
	DateSent    time.Time `json:"date_sent"`
	SenderID    uint8     `json:"sender_id"`
	Sender      User      `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"user"`
	RecipientID uint8
	Recipient   User `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"recipient"`
	CalenderID  uint8
	Calender    Calender `gorm:"constraint:onUpdate:CASCADE,onDelete: SET NULL;" json:"calender""`
}

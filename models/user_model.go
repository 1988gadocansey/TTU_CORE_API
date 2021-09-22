package models

import "gorm.io/gorm"

type User struct {
	Base
	gorm.Model
	Name       string `json:"Name"`
	Email      string `db:"Email" gorm:"unique" json:"Email"`
	Password   string `db:"Passowrd" json:"-"`
	IsVerified bool   `db:"IsVerified" json:"IsVerified"`
	Level      int    `db:"Level" json:"level"`
	Status     int    ` db:"Status" json:"Status"`
	Role       byte   `db:"Role" json:"Role"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("Role", "User")
	}
	return
}

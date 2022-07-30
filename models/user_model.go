package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	gorm.Model
	Name       string `json:"Name"`
	Email      string `db:"Email" gorm:"unique" json:"Email"`
	Password   []byte `json:"_" validate:"required,min=8"`
	IsVerified bool   `db:"IsVerified" json:"IsVerified"`
	Level      int    `db:"Level" json:"level"`
	Status     int    ` db:"Status" json:"Status"`
	//Role       byte   `db:"Role" json:"Role"`
	RoleId uint `json:"role_id"`
	Role   Role `json:"role" gorm:"foreignKey:role_id"`
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {

	tx.Model(user).Update("Role", 1)

	return
}
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = []byte(string(hashedPassword))
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&products)

	return products
}

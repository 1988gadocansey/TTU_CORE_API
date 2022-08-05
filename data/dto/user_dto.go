package dto

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email,min=5"`
	Password []byte `json:"password" validate:"required,min=8"`
	//Password  []byte `json:"-"`
	IsVerified bool `json:"is_verified"`
	Level      int  `json:"level"  validate:"required,number"`
	Status     int  `json:"status" validate:"required,number"`
	BaseDto
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

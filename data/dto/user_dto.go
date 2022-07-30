package dto

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

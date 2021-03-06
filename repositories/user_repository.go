package repositories

import (
	"TTU_CORE_ERP_API/data/dto"
	"TTU_CORE_ERP_API/models"
	"fmt"
)

func CreateUser(userData dto.User) string {
	err := DB.Create(userData).Error
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return "OK"
}

func GetUserByEmail(email string) models.User {
	var user_data models.User
	err := DB.Where("email = ?", email).First(&user_data)
	if err != nil {
		fmt.Println(err)
	}
	return user_data
}

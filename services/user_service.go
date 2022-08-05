package services

import (
	"TTU_CORE_ERP_API/data/dto"
	"TTU_CORE_ERP_API/data/response"
	"TTU_CORE_ERP_API/repositories"
	"TTU_CORE_ERP_API/validations"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func CreateUser(user *dto.User) response.DataResponse {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	baseData := dto.BaseDto{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userData := dto.User{
		Name:       user.Name,
		Email:      user.Email,
		Password:   password,
		IsVerified: user.IsVerified,
		Level:      user.Level,
		Status:     user.Status,
		BaseDto:    baseData,
	}

	// Set Validation Struct
	validations.UserStruct = userData
	// Validate
	errors := validations.ValidateUserStruct()
	if errors != nil {
		response := response.DataResponse{
			Status:  fiber.StatusNotAcceptable,
			Message: "Input error",
			Error:   errors,
		}
		return response
	}

	// Handle repositories error
	repositoryResponse := repositories.CreateUser(userData)
	if repositoryResponse != "OK" {
		response := response.DataResponse{
			Status:  1062,
			Message: repositoryResponse,
		}
		return response
	}

	data, _ := json.Marshal(userData)

	response := response.DataResponse{
		Status:  fiber.StatusOK,
		Message: "Creation successful",
		Data:    string(data),
	}
	return response
}

func AuthenticateUser(credentials map[string]string) response.DataResponse {
	userData := repositories.GetUserByEmail(credentials["email"])
	if userData.ID == 0 {
		response := response.DataResponse{
			Status:  fiber.StatusNotFound,
			Message: "Account not found.",
		}
		return response
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(credentials["password"])); err != nil {
		response := response.DataResponse{
			Status:  fiber.StatusNotFound,
			Message: "Incorrect password.",
		}
		return response
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userData.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		response := response.DataResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Cannot generate token.",
		}
		fmt.Println(err)
		return response
	}

	response := response.DataResponse{
		Status:  fiber.StatusOK,
		Message: "Authentication successful.",
		Data:    token,
	}
	return response
}

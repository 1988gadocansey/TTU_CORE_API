package controllers

import (
	"TTU_CORE_ERP_API/data/dto"
	"TTU_CORE_ERP_API/data/response"
	"TTU_CORE_ERP_API/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := new(dto.User)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	response := services.CreateUser(user)
	return ctx.JSON(response)
}

func AuthenticateUser(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	response := services.AuthenticateUser(data)
	return ctx.JSON(response)
}

func UserList(ctx *fiber.Ctx) error {
	response := response.DataResponse{
		Status:  200,
		Message: ctx.Params("id"),
	}
	return ctx.JSON(response)
}

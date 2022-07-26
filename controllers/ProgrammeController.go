package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var programmeRepository repositories.ProgrammeRepository

func init() {
	programmeRepository = repositories.NewProgrammeRepository()
}
func AllProgramme(c *fiber.Ctx) error {
	programme := programmeRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    programme,
		Title:   "GetAllProgramme",
		Message: "All programmes",
	}

	return c.Status(resp.Code).JSON(resp)
}
func GetProgramme(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting programme information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	programmes, err := programmeRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting programme information",
		}
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    programmes,
		Title:   "GetSingleProgramme",
		Message: "programme data retrieved successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

/*
	AddNewProgramme adds new programme object
*/
func AddNewProgramme(c *fiber.Ctx) error {
	programme := &models.Programme{}

	err := c.BodyParser(programme)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing programme body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := programmeRepository.Save(*programme)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new programme",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	programme, err = programmeRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added programme",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if programme == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("programme with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding programme",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    programme,
		Title:   "OK",
		Message: "new programme added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var facultyRepository repositories.FacultyRepository

func init() {
	facultyRepository = repositories.NewFacultyRepository()
}
func GetAllFaculties(c *fiber.Ctx) error {
	faculties := facultyRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    faculties,
		Title:   "GetAllFaculties",
		Message: "All Faculties",
	}

	return c.Status(resp.Code).JSON(resp)
}
func GetFaculty(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting faculty information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	faculties, err := facultyRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting faculty information",
		}
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    faculties,
		Title:   "GetSingleFaculty",
		Message: "faculty data retrieved successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

/*
	AddNewFaculty adds new faculty object
*/
func AddNewFaculty(c *fiber.Ctx) error {
	faculty := &models.Faculty{}

	err := c.BodyParser(faculty)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing faculty body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := facultyRepository.Save(*faculty)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new faculty",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	faculty, err = facultyRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added faculty",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if faculty == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("faculty with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding faculty",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    faculty,
		Title:   "OK",
		Message: "new faculty added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

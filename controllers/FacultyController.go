package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"github.com/gofiber/fiber/v2"
	"net/http"
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

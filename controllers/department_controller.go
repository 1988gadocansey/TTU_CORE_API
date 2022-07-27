package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var departmentRepository repositories.DepartmentRepository

func init() {
	departmentRepository = repositories.NewDepartmentRepository()
}
func GetAllDepartments(c *fiber.Ctx) error {
	departments := departmentRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    departments,
		Title:   "GetAllDepartments",
		Message: "All Departments",
	}

	return c.Status(resp.Code).JSON(resp)
}
func GetDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting department information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	faculties, err := departmentRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting department information",
		}
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    faculties,
		Title:   "GetSingleDepartment",
		Message: "department data retrieved successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

/*
	AddNewDepartment adds new faculty object
*/
func AddNewDepartment(c *fiber.Ctx) error {
	department := &models.Department{}

	err := c.BodyParser(department)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing department body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := departmentRepository.Save(*department)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new department",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	department, err = departmentRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added department",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if department == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("department with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding department",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    department,
		Title:   "OK",
		Message: "new department added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

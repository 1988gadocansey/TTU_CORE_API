package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var studentRepository repositories.StudentRepository

func init() {
	studentRepository = repositories.NewStudentRepository()
}

// GetAllLStudent db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key
//db.Model(&user).Related(&languages, "Languages")
//// SELECT * FROM "languages" INNER JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" WHERE "user_languages"."user_id" = 111
func GetAllLStudent(c *fiber.Ctx) error {
	levels := levelRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    levels,
		Title:   "GetAllLevel",
		Message: "All Level",
	}

	return c.Status(resp.Code).JSON(resp)
}

func GetSingleStudent(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting level information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)

	}

	level, err := levelRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting level information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if level == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("level with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding level",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    level,
		Title:   "OK",
		Message: "Level information",
	}
	return c.Status(resp.Code).JSON(resp)

}

// AddNewStudent adds new level
func AddNewStudent(c *fiber.Ctx) error {
	// lets say we are using request dto
	//students := &request.StudentRequest{}

	student := &models.Student{}

	err := c.BodyParser(student)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing student body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := studentRepository.Save(*student)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new level",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	student, err = studentRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added student",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if student == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("level with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding student",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    student,
		Title:   "OK",
		Message: "new student added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

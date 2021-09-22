package controllers

import (
	"TTU_CORE_PORTAL_GO/models"
	"TTU_CORE_PORTAL_GO/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

var levelRepository repositories.LevelRepository

func init() {
	levelRepository = repositories.NewLevelRepository()
}

// GetAllLevel GetAllLevels gets all repository information
func GetAllLevel(c *fiber.Ctx) error {
	levels := levelRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    levels,
		Title:   "GetAllLevel",
		Message: "All Level",
	}

	return c.Status(resp.Code).JSON(resp)
}

// GetSingleCompany Gets single company information
func GetSingleLevel(c *fiber.Ctx) error {
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

// AddNewLevel adds new level
func AddNewLevel(c *fiber.Ctx) error {
	level := &models.Level{}

	err := c.BodyParser(level)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing level body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := levelRepository.Save(*level)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new level",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	level, err = levelRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added level",
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
		Message: "new level added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

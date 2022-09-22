package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var productRepository repositories.ProductRepository

func init() {
	productRepository = repositories.NewProductRepository()
}

func GetAllProduct(c *fiber.Ctx) error {
	products := productRepository.FindAll()

	resp := models.Response{
		Code:    http.StatusOK,
		Body:    products,
		Title:   "GetAllProducts",
		Message: "All Products",
	}

	return c.Status(resp.Code).JSON(resp)
}
func FindByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	product, err := productRepository.FindByCode(code)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in fetching product information",
		}
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    product,
		Title:   "Product by code",
		Message: "Getting product by",
	}
	return c.Status(resp.Code).JSON(resp)
}
func CreateProduct(c *fiber.Ctx) error {
	data := &models.PaymentProduct{}

	err := c.BodyParser(data)

	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing product body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	_, err = productRepository.Save(*data)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new product",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    data,
		Title:   "OK",
		Message: "new product added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

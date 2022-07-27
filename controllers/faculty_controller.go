package controllers

import (
	"TTU_CORE_ERP_API/models"
	"TTU_CORE_ERP_API/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	"log"
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

/*
	This function export faculty data to excel
*/
func ExportToExcel(c *fiber.Ctx) error {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	err := f.SetCellValue("Sheet2", "A2", "Hello world.")
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "B2", 100)
	if err != nil {
		return err
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return nil
}
func ExportChart(c *fiber.Ctx) error {

	categories := map[string]string{"A1": "USA", "A2": "China", "A3": "UK",
		"A4": "Russia", "A5": "South Korea", "A6": "Germany"}

	values := map[string]int{"B1": 46, "B2": 38, "B3": 29, "B4": 22, "B5": 13, "B6": 11}

	f := excelize.NewFile()

	for k, v := range categories {

		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}

	for k, v := range values {

		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}

	if err := f.AddChart("Sheet1", "E1", `{
        "type":"col", 
        "series":[
            {"name":"Sheet1!$A$2","categories":"Sheet1!$A$1:$A$6",
                "values":"Sheet1!$B$1:$B$6"}
            ],
            "title":{"name":"Olympic Gold medals in London 2012"}}`); err != nil {

		log.Fatal(err)
	}

	if err := f.SaveAs("gold_medals.xlsx"); err != nil {
		log.Fatal(err)
	}
	return nil
}
func ExportChartMain(c *fiber.Ctx) error {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range values {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return nil
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return nil
}

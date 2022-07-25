package routes

import (
	"TTU_CORE_ERP_API/controllers"
	"TTU_CORE_ERP_API/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markbates/pkger"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	/*
	 * Version 1
	 */
	v1 := api.Group("/v1", middlewares.RequestAuthencation)

	user := v1.Group("/user")
	user.Post("/create", controllers.CreateUser)
	user.Post("/login", controllers.AuthenticateUser)
	user.Get("/:id", controllers.UserList)

	// Group Level related APIs
	levelGroup := v1.Group("/levels")
	levelGroup.Get("/", controllers.GetAllLevel)
	levelGroup.Get("/:id", controllers.GetSingleLevel)
	levelGroup.Post("/", controllers.AddNewLevel)

	// Group Faculty related APIs
	facultyGroup := v1.Group("/faculties")
	facultyGroup.Get("/", controllers.GetAllFaculties)
	facultyGroup.Get("/:id", controllers.GetFaculty)
	facultyGroup.Post("/", controllers.AddNewFaculty)

	// Group Department related APIs
	departmentGroup := v1.Group("/departments")
	departmentGroup.Get("/", controllers.GetAllDepartments)
	departmentGroup.Get("/:id", controllers.GetDepartment)
	departmentGroup.Post("/", controllers.AddNewDepartment)

	// Group Student related APIs
	studentGroup := v1.Group("/students")
	studentGroup.Get("/", controllers.GetAllLStudent)
	studentGroup.Get("/:id", controllers.GetSingleStudent)
	studentGroup.Post("/", controllers.AddNewStudent)

	cdn := app.Group("/cdn")
	cdn.Use("/images", filesystem.New(
		filesystem.Config{
			Root: pkger.Dir("/public/images"),
		},
	))
}

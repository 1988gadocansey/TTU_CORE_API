package main

import (
	"TTU_CORE_ERP_API/configs"
	"TTU_CORE_ERP_API/repositories"
	"TTU_CORE_ERP_API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	db := database.Init()
	/*err := database.Migrate()
	if err != nil {
		return
	}*/
	repositories.DB = db
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	// Use logger
	app.Use(logger.New())

	routes.Setup(app)

	app.Listen(":5000")

}

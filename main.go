package main

import (
	"TTU_CORE_PORTAL_GO/configs"
	"TTU_CORE_PORTAL_GO/repositories"
	"TTU_CORE_PORTAL_GO/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/hojabri/backend/repository"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	db := database.Init()
	err := database.Migrate()
	if err != nil {
		return
	}
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

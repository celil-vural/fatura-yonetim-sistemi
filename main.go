package main

import (
	"fatura-yonetim-sistemi/handler"
	"fatura-yonetim-sistemi/infrastructure/database"
	"fatura-yonetim-sistemi/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	log.Fatal(app.Listen(":3030"))
}

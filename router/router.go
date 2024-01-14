package router

import (
	"fatura-yonetim-sistemi/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/login", handler.Login)
}

package router

import (
	"fatura-yonetim-sistemi/handler"
	"fatura-yonetim-sistemi/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/login", handler.Login)
	manager := v1.Group("/manager", middleware.AddJwtMiddleware(), middleware.IsManager())
	manager.Post("/addUser", handler.Register)
	manager.Post("/addApartment", handler.CreateApartment)
}

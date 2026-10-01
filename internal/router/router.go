package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
)

func SetupRoutes(fiberApp *fiber.App, application *app.App) {
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Ticket Booking APP 🏃")
	})

	api := fiberApp.Group("/api/v1")

	// ==================| User Routes |==================
	api.Post("/users", application.Models.UserHandler.CreateUser)
}

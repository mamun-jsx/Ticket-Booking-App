package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/booking"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user"
)

// SetupRoutes initializes all global middleware and module routes for the application.
func SetupRoutes(fiberApp *fiber.App, application *app.App) {
	// Root health check endpoint
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Ticket Booking APP 🏃")
	})

	// API version 1 route group
	api := fiberApp.Group("/api/v1")

	// ==================| Module Routes Register |==================
	// Register user domain endpoints (e.g. POST /api/v1/users/create)

	user.InitUserRoutes(api, application.DB)
	event.InitEventRoutes(api, application.DB)
	booking.InitBookingRoutes(api, application.DB)
}

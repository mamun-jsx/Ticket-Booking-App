package booking

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event"
	"gorm.io/gorm"
)

// InitEventRoutes initializes repository, service, handler and registers event routes.
func InitBookingRoutes(router fiber.Router, db *gorm.DB) {
	bookingRepo := NewBookingRepository(db)
	eventRepo := event.NewRepository(db)
	bookingServ := NewBookingService(bookingRepo, eventRepo)
	bookingHandler := NewBookingHandler(bookingServ)

	api := router.Group("/booking")

	api.Post("/create", bookingHandler.CreateBooking)
}

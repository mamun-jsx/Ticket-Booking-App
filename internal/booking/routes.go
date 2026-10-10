package booking

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/auth"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/middlewares"
	"gorm.io/gorm"
)

// InitBookingRoutes initializes repository, service, handler and registers booking routes.
func InitBookingRoutes(router fiber.Router, db *gorm.DB) {
	bookingRepo := NewBookingRepository(db)
	eventRepo := event.NewRepository(db)
	bookingServ := NewBookingService(bookingRepo, eventRepo)
	bookingHandler := NewBookingHandler(bookingServ)

	jwtService := auth.NewJWTService("jwt_sercret_key", 24*time.Hour)

	api := router.Group("/booking", middlewares.AuthMiddleware(jwtService))

	api.Post("/create", bookingHandler.CreateBooking)
}


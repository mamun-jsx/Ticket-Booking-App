package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/router"
)

// HTTPServer wraps the Fiber application instance and configuration.
type HTTPServer struct {
	app  *fiber.App
	port string
}

// NewHTTPServer initializes a new HTTP server with configured routes and middleware.
func NewHTTPServer(application *app.App) *HTTPServer {
	fiberApp := fiber.New()

	// Register routes
	router.SetupRoutes(fiberApp, application)

	return &HTTPServer{
		app:  fiberApp,
		port: application.Config.AppPort,
	}
}

// Start begins listening for incoming HTTP requests.
func (s *HTTPServer) Start() error {
	address := fmt.Sprintf(":%s", s.port)
	log.Printf("🚀 Server is running on http://localhost%s", address)
	return s.app.Listen(address)
}
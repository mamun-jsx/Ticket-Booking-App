package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/router"
)

func main() {
	// 1. Initialize application configuration and database connection
	application := app.BootStrapApp()

	// 2. Initialize a new Fiber web application instance
	fiberApp := fiber.New()

	// 3. Register all API endpoints and application routes
	router.SetupRoutes(fiberApp, application)

	// 4. Construct server listening address using the port from configuration
	serverAddress := fmt.Sprintf(":%s", application.Config.AppPort)

	// 5. Start the HTTP server
	log.Printf("🚀 Server is running on http://localhost%s", serverAddress)
	if err := fiberApp.Listen(serverAddress); err != nil {
		log.Fatalf("server failed to run: %v", err)
	}
}

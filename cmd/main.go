package main

import (
	"log"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/server"
)

func main() {
	// 1. Initialize application configuration and database connection
	application := app.BootStrapApp()

	// 2. Initialize HTTP server
	httpServer := server.NewHTTPServer(application)

	// 3. Start the HTTP server
	if err := httpServer.Start(); err != nil {
		log.Fatalf("server failed to run: %v", err)
	}
}


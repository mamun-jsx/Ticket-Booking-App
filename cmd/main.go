package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/app"
)

func main() { 
	application := app.BootStrapApp() // call app file
	fiberApp := fiber.New() // make fiber app 

	// make server address
	serverAddress := fmt.Sprintf(":%s", application.Config.AppPort)

	// Server listen on port
	log.Printf("🚀 Server is running on http://localhost%s", serverAddress)
	if err := fiberApp.Listen(serverAddress); err != nil {
		log.Fatalf("server failed to run: %v", err)
	}
}

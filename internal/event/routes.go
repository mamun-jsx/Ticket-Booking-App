package event

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// InitEventRoutes initializes repository, service, handler and registers event routes.
func InitEventRoutes(router fiber.Router, db *gorm.DB) {
	eventRepo := NewRepository(db)
	eService := NewService(eventRepo)
	eventHandler := NewHandler(eService)

	api := router.Group("/event")

	api.Get("", eventHandler.GetAllEvents)
	api.Get("/:id", eventHandler.GetEventByID)
	api.Post("/create", eventHandler.CreateEvent)
	api.Put("/update/:id", eventHandler.UpdateEvent)
	api.Delete("/delete/:id", eventHandler.DeleteEvent)
}
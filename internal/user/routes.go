package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// InitUserRoutes initializes repository, service, handler and registers user routes.
func InitUserRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	users := router.Group("/users")
	users.Post("/create", userHandler.CreateUser)
}

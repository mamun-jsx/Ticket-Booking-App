package user

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/auth"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/middlewares"
	"gorm.io/gorm"
)

// InitUserRoutes initializes repository, service, handler and registers user routes.
func InitUserRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := NewRepository(db)
	jwtService := auth.NewJWTService("jwt_sercret_key", 24*time.Hour)
	userService := NewService(userRepo, jwtService)
	userHandler := NewHandler(userService)

	users := router.Group("/auth")
	users.Post("/register", userHandler.CreateUser)
	users.Post("/login", userHandler.LoginUser)
	users.Get("/me", middlewares.AuthMiddleware(jwtService), userHandler.AuthUser)
}

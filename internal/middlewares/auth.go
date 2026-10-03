package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/auth"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/httpresponse"
)

// AuthMiddleware validates the JWT token from the Authorization header.
func AuthMiddleware(jwtService auth.JWTService) fiber.Handler {
	
	
	return func(c fiber.Ctx) error {
		// get auth from header
		authHeader := c.Get("Authorization") 

		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Code:    http.StatusUnauthorized,
				Message: "Authorization header is missing",
			})
		}

		// check bearer token
		parse := strings.Split(authHeader, " ")
		if len(parse) != 2 || parse[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Code:    http.StatusUnauthorized,
				Details: "Missing Bearer Token",
			})
		}

		// validate token
		tokenString := parse[1]

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Code:    http.StatusUnauthorized,
				Message: "Invalid or expired token",
			})
		}

		// store user into locals
		c.Locals("user_id", claims.UserID)   
		c.Locals("user_email", claims.Email)
		c.Locals("user_name", claims.Name)
		c.Locals("user", claims)

		return c.Next() // call next handler
	} 
} 

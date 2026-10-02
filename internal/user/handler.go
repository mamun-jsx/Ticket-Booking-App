package user

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/httpresponse"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user/dto"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateUser handles HTTP request to create a new user.
func (h *Handler) CreateUser(c fiber.Ctx) error {
	var req dto.CreateRequest // user input

	// 1. Bind and parse JSON request body into the CreateRequest DTO
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body Input",
			Details: err.Error(),
		})

	}

	// 2. Delegate business logic and entity creation to the service layer
	res, err := h.service.CreateUser(&req)

	// 3. Return an error response if service layer returns an error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create user",
			Details: err.Error(),
		})
	}

	// 4. Return HTTP 201 Created with the user response DTO on success
	return c.Status(http.StatusCreated).JSON(res)
}

// LoginUser handles user authentication request.
func (h *Handler) LoginUser(c fiber.Ctx) error {
	var req dto.LoginRequest

	// 1. Bind and parse JSON request body
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body Input",
			Details: err.Error(),
		})
	}

	// 2. Authenticate user via service layer
	res, err := h.service.LoginUser(&req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed",
			Details: err.Error(),
		})
	}

	// 3. Return HTTP 200 OK with user response DTO
	return c.Status(http.StatusOK).JSON(res)
}

package user

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/httpresponse"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user/dto"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

// CreateUser handles HTTP request to create a new user.
func (h *handler) CreateUser(c fiber.Ctx) error {
	var req dto.CreateRequest // user input

	// ১. ক্লায়েন্ট থেকে আসা রিকোয়েস্ট বডি পার্স/বাইন্ড করা হচ্ছে
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body Input",
			Details: err.Error(),
		})

	}

	// ২. সার্ভিস লেয়ার কল করে বিজনেস লজিক সম্পন্ন করা হচ্ছে
	res, err := h.service.CreateUser(&req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create user",
			Details: err.Error(),
		})
	}

	// ৩. সফল হলে স্ট্যাটাস 201 Created সহ JSON রেসপন্স রিটার্ন করা হচ্ছে
	return c.Status(http.StatusCreated).JSON(res)
}


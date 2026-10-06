package booking

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/booking/dto"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/httpresponse"
)

type handler struct {
	service BookingService
}

func NewBookingHandler(s BookingService) *handler {
	return &handler{service: s}
}


// ? func to get current user.

func getCurrentUserID(c fiber.Ctx) (uint, bool) {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return 0, false
	}
	return userID, true
}
func bookingErrorResponse(c fiber.Ctx, err error) error {
	if errors.Is(err, ErrBookingNotFound) {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Booking Not Found",
			Details: err.Error(),
		})
	}
	if errors.Is(err, event.ErrEventNotFount) {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Event Not Found",
			Details: err.Error(),
		})
	}
	return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server error",
		Details: err.Error(),
	})
}
func (h *handler) CreateBooking(c fiber.Ctx) error {
	userId, ok := getCurrentUserID(c)

	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorize",
		})
	}

	var req dto.CreateBookingRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
	}
	response, err := h.service.CreateBooking(userId, &req)
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.Status(http.StatusCreated).JSON(response)
}

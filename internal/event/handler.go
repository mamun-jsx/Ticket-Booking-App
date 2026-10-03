package event

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event/dto"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/httpresponse"
)

var ErrEventNotFount = errors.New("Event Not found")

type handler struct {
	service EventService
}

func NewHandler(s EventService) *handler {
	return &handler{service: s}
}

// error handling
func eventErrorResponse(c fiber.Ctx, err error) error {
	if errors.Is(err, ErrEventNotFount) {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Event not found",
			Details: err.Error(),
		})
	}
	return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "somthing went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateEvent(c fiber.Ctx) error {
	var req dto.CreateRequestEvent
	// 1. Bind and parse JSON request body into the CreateRequest DTO
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body Input",
			Details: err.Error(),
		})
	}
	// 2. Delegate business logic and entity creation to the service layer
	res, err := h.service.CreateEvent(&req)

	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.Status(http.StatusOK).JSON(res)
}

// get the array of object as (all events)
func (h *handler) GetAllEvents(c fiber.Ctx) error {
	events, err := h.service.GetAllEvents()
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.Status(http.StatusOK).JSON(events)
}

// * Get a single Event via ID
func (h *handler) GetEventByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Event ID",
			Details: err.Error(),
		})
	}
	
	res, err := h.service.GetEventByID(uint(id))
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.Status(http.StatusOK).JSON(res)
}

package booking

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/booking/dto"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/event"
)

var (
	ErrBookingNotFounded      = errors.New("booking not found")
	ErrNotEnoughTickets       = errors.New("not enough tickets")
	ErrBookingAlreadyCanceled = errors.New("booking already canceled")
	ErrForbiddenBookingAccess = errors.New("you are not authorized to do that")
)

type BookingService interface {
	CreateBooking(userID uint, req *dto.CreateBookingRequest) (*dto.BookingResponse, error)
	GetUserBookings(userID uint) ([]*dto.BookingResponse, error)
	CancelBooking(bookingID uint, userID uint) error
}

type service struct {
	bookingRepo BookingRepository
	eventRepo   event.EventRepository
}

func NewBookingService(bookingRepo BookingRepository, eventRepo event.EventRepository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

// create a single booking

func generateBookingCode() string {
	return "GT-" + uuid.New().String()[:8]
}

func (s *service) CreateBooking(userId uint, req *dto.CreateBookingRequest) (*dto.BookingResponse, error) {
	event, err := s.eventRepo.GetEventByID(req.EventID) // get event id 
	if err != nil {
		return nil, err
	}

	// check tickets available or not
	if event.AvailableTickets < req.Quantity {
		return nil, ErrBookingNotFounded
	}
	// create a single booking
	booking := &Booking{
		UserID:      userId,
		EventID:     req.EventID,
		Quantity:    req.Quantity,
		Status:      BookingConfirmed,
		TotalPrice:  req.Quantity * event.Price,
		BookingCode: generateBookingCode(),
	}
	if err := s.bookingRepo.Create(booking); err != nil {
		return nil, err
	}

	// reduce the event tickets
	event.AvailableTickets = event.AvailableTickets - req.Quantity
	if err := s.eventRepo.UpdateEvent(event); err != nil {
		return nil, err
		// ? Update database and the event tickets quantity
	}
	return booking.ToResponse(), nil
}

// get all bookings for a user
func (s *service) GetUserBookings(userID uint) ([]*dto.BookingResponse, error) {
	bookings, err := s.bookingRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.BookingResponse
	for _, b := range bookings {
		responses = append(responses, b.ToResponse())
	}
	return responses, nil
}

// cancel a booking by ID (only the owner can cancel)
func (s *service) CancelBooking(bookingID uint, userID uint) error {
	booking, err := s.bookingRepo.GetBookingById(bookingID)
	if err != nil {
		return err
	}

	// ensure the booking belongs to the requesting user
	if booking.UserID != userID {
		return ErrForbiddenBookingAccess
	}

	// prevent double-cancellation
	if booking.Status == BookingCanclled {
		return ErrBookingAlreadyCanceled
	}

	// restore available tickets on the event
	event, err := s.eventRepo.GetEventByID(booking.EventID)
	if err != nil {
		return err
	}
	event.AvailableTickets += booking.Quantity
	if err := s.eventRepo.UpdateEvent(event); err != nil {
		return err
	}

	// mark the booking as cancelled
	booking.Status = BookingCanclled
	return s.bookingRepo.UpdateBooking(booking)
}

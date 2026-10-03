package event

import (
	"time"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/event/dto"
)

type EventService interface {
	CreateEvent(event *Event) error
	GetAllEvents() ([]*Event, error)
	GetEventByID(id uint64) (*Event, error)
	UpdateEvent(event *Event) error
	DeleteEvent(event *Event) error
}

type service struct {
	repo EventRepository
}

func NewService(repo EventRepository) EventService {
	// Return a pointer to the service struct
	return &service{repo: repo}
}

func (s *service) CreateEvent(req *dto.CreateRequestEvent) (*dto.ResponseEvent, error) {
	
	parsedTime, err := time.Parse(time.RFC3339, req.StartsAt)
	if err != nil {
		return nil, err // Return parsing errors (e.g., bad format from client)
	}

	// convert dto into event model
	event := Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartsAt:         parsedTime,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}

	if err := s.repo.CreateEvent(&event); err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

func (s *service) GetAllEvents() ([]*Event, error) {
	return s.repo.GetAllEvents()
}

func (s *service) GetEventByID(id uint64) (*Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *service) UpdateEvent(event *Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *service) DeleteEvent(event *Event) error {
	return s.repo.DeleteEvent(event)
}

package event

import (
	"time"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/event/dto"
)

type EventService interface {
	CreateEvent(req *dto.CreateRequestEvent) (*dto.ResponseEvent, error)
	GetAllEvents() ([]*dto.ResponseEvent, error)
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

// create an Evenet
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

func (s *service) GetAllEvents() ([]*dto.ResponseEvent, error) {
	events, err := s.repo.GetAllEvent()

	if err != nil {
		return nil, err
	}
	var allEvents []*dto.ResponseEvent
	for _, event := range events {
		allEvents = append(allEvents, event.ToResponse())
	}
	// return te response
	return allEvents, nil
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

package event

import (
	"errors"
	"time"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/event/dto"
)

type EventService interface {
	CreateEvent(req *dto.CreateRequestEvent) (*dto.ResponseEvent, error)
	GetAllEvents() ([]*dto.ResponseEvent, error)
	GetEventByID(id uint) (*dto.ResponseEvent, error)
	UpdateEvent(eventId uint, req *dto.UpdateRequestEvent) (*dto.ResponseEvent, error)
	DeleteEvent(id uint) error
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

// ? All Events []
func (s *service) GetAllEvents() ([]*dto.ResponseEvent, error) {
	events, err := s.repo.GetAllEvent()

	if err != nil {
		return nil, err
	}
	var allEvents []*dto.ResponseEvent
	for _, event := range events {
		allEvents = append(allEvents, event.ToResponse())
	}
	if len(allEvents) == 0 {
		return nil, errors.New("Currently no Events posted")
	}
	// return te response
	return allEvents, nil
}

// * get a single events By ID
func (s *service) GetEventByID(id uint) (*dto.ResponseEvent, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

// ? Update a single service
func (s *service) UpdateEvent(eventId uint, req *dto.UpdateRequestEvent) (*dto.ResponseEvent, error) {
	event, err := s.repo.GetEventByID(eventId)
	if err != nil {
		return nil, err
	}
	if req.Title != "" {
		event.Title = req.Title
	}
	if req.Description != "" {
		event.Description = req.Description
	}
	if req.Location != "" {
		event.Location = req.Location
	}
	if req.StartsAt != "" {
		parsedTime, err := time.Parse(time.RFC3339, req.StartsAt)
		if err != nil {
			return nil, err // Return parsing errors (e.g., bad format from client)
		}
		event.StartsAt = parsedTime
	}
	if req.TotalTickets != 0 {
		event.TotalTickets = req.TotalTickets
	}
	if req.Price != 0 {
		event.Price = req.Price
	}

	if err := s.repo.UpdateEvent(event); err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

// ! delete a single event by ID
func (s *service) DeleteEvent(id uint) error {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteEvent(event)
}
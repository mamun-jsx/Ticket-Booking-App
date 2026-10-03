package event

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

func (s *service) CreateEvent(event *Event) error {
	return s.repo.CreateEvent(event)
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

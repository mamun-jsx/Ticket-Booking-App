package event

import "gorm.io/gorm"

type EventRepository interface {
	CreateEvent(event *Event) error
	GetAllEvents() ([]*Event, error)
	GetEventByID(id uint64) (*Event, error)
	UpdateEvent(event *Event) error
	DeleteEvent(event *Event) error
}

type repository struct {
	db gorm.DB
}

// make an instance of repository
func NewRepository(db *gorm.DB) EventRepository {
	return &repository{db: *db}
}

// create an event
func (r *repository) CreateEvent(event *Event) error {
	return r.db.Create(event).Error
}

// get all events
func (r *repository) GetAllEvents() ([]*Event, error) {
	var events []*Event
	err := r.db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

// get an event by id
func (r *repository) GetEventByID(id uint64) (*Event, error) {
	var event *Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

// update an event
func (r *repository) UpdateEvent(event *Event) error {
	err := r.db.Save(event).Error
	if err != nil {
		return err
	}
	return nil
}

// delete an event
func (r *repository) DeleteEvent(event *Event) error {
	err := r.db.Delete(event).Error
	if err != nil {
		return err
	}
	return nil
}
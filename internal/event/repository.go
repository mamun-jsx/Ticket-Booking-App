package event

import "gorm.io/gorm"

type Repository interface {
	CreateEvent(event *Event) error
	GetAllEvents() ([]*Event, error)
	GetEventByID(id uint64) (*Event, error)
	UpdateEvent(event *Event) error
	DeleteEvent(event *Event) error
}
type repository struct {
	db gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: *db}
}

func (r *repository) CreateEvent(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) GetAllEvents() ([]*Event, error) {
	var events []*Event
	return events, r.db.Find(&events).Error
}

func (r *repository) GetEventByID(id uint64) (*Event, error) {
	var event *Event
	return event, r.db.First(&event, id).Error
}

func (r *repository) UpdateEvent(event *Event) error {
	return r.db.Save(event).Error
}

func (r *repository) DeleteEvent(event *Event) error {
	return r.db.Delete(event).Error
}
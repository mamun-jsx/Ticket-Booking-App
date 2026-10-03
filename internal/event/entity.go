package event

import (
	"time"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/event/dto"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(150);not null"`
	Description string `json:"description" gorm:"type:text"`
	Location    string `json:"location" gorm:"type:varchar(255);not null"`

	StartsAt         time.Time `json:"start_at"`
	TotalTickets     int       `json:"total_tickets" gorm:"not null;check:(total_tickets > 0)"`
	AvailableTickets int       `json:"available_tickets" gorm:"not null"`
	Price            int       `json:"price" gorm:"not null;check:(price > 0)"`
}

func (e *Event) ToResponse() *dto.ResponseEvent {
	return &dto.ResponseEvent{
		ID:               uint64(e.ID),
		Title:            e.Title,
		Description:      e.Description,
		Location:         e.Location,
		StartsAt:         e.StartsAt,
		TotalTickets:     e.TotalTickets,
		AvailableTickets: e.AvailableTickets,
		Price:            e.Price,
	}
}

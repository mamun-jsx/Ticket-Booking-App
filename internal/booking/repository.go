package booking

import "gorm.io/gorm"

type BookingRepository interface {
	Create(booking *Booking) error
	GetBookingById(bookingId uint) (*Booking, error)
	GetByUserID(userID uint) ([]*Booking, error)
	UpdateBooking(booking *Booking) error
}

type repository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &repository{db: db}
}

func (r *repository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}
func (r *repository) GetBookingById(bookingId uint) (*Booking, error) {
	return nil, nil
}

func (r *repository) GetByUserID(userID uint) ([]*Booking, error) {
	return nil, nil
}

func (r *repository) UpdateBooking(booking *Booking) error {
	return r.db.Save(booking).Error
}

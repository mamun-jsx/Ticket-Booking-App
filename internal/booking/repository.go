package booking

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorBookingNotFound = errors.New("booking not found")
var ErrorBookingAlreadyCancelled = errors.New("booking already cancelled")

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

// create a single booking
func (r *repository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}

// get a single booking by booking id
func (r *repository) GetBookingById(bookingId uint) (*Booking, error) {
	var booking Booking
	err := r.db.First(&booking, bookingId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorBookingNotFound
		}
		return nil, err
	}
	return &booking, nil
}

// get booking history by User id
func (r *repository) GetByUserID(userID uint) ([]*Booking, error) {
	var bookings []*Booking
	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *repository) UpdateBooking(booking *Booking) error {
	return r.db.Save(booking).Error
}

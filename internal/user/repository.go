package user

import (
	"errors"

	"gorm.io/gorm"
)

// ErrorAlreadyExist is returned when attempting to create a user with an email that already exists.
var ErrorAlreadyExist = errors.New("user with this email already exist")

// UserRepository defines the contract for user database operations.
type UserRepository interface {
	CreateUser(user *User) error
}

// repository implements UserRepository interface with a GORM database instance.
type repository struct {
	db *gorm.DB
}

// NewRepository creates and returns a new UserRepository implementation.
func NewRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

// CreateUser inserts a new user record into the PostgreSQL database.
func (r *repository) CreateUser(user *User) error {
	// Execute INSERT query via GORM
	result := r.db.Create(&user)
	if result.Error != nil {
		// Detect unique constraint violations (e.g., duplicated email)
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}
		return result.Error
	}
	return nil
}

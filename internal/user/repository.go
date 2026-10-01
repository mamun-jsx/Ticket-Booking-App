package user

import (
	"errors"

	"gorm.io/gorm"
)

// ? database related operation and interaction into here
var ErrorAlreadyExist = errors.New("user with this email already exist")

type UserRepository interface {
	CreateUser(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}
		return result.Error
	}
	return nil
}

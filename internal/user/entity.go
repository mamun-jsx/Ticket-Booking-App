package user

import "gorm.io/gorm"

// "time"

// "github.com/google/uuid"
// "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

package user

import "gorm.io/gorm"

// User represents the database entity for users in PostgreSQL.
// It embeds gorm.Model which provides ID, CreatedAt, UpdatedAt, and DeletedAt fields.
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

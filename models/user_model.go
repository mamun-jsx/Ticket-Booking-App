package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Name     string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email    string `gorm:"type:varchar(255);not null" json:"name"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}


package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/Ticket-Booking-App/config"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DbUrl == "" {
		return nil, fmt.Errorf("Database URL Not FOUND")
	}

	// open a connection request to postgresql via Gorm
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DbUrl,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("Database connection failed: %v", err)
		return nil, err
	}

	// Auto-migrate domain entities directly
	err = db.AutoMigrate(
		&user.User{},
		// Future domain entities will be added here, e.g.: &ticket.Ticket{}
	)
	if err != nil {
		log.Printf("Database migration failed: %v", err)
		return nil, err
	}
	log.Println("🚀🚀🚀🚀 PostgreSQL Migration successful 🚀🚀🚀🚀")
	return db, nil
}

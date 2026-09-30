package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/Ticket-Booking-App/config"
	"github.com/mamun-jsx/Ticket-Booking-App/models"
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
	// Unpack the slice automatically via the models func from model package
	// get the user model from dynamic func GetRegisterModels and auto migrate to database
	err = db.AutoMigrate(models.GetRegisterModels()...)
	if err != nil {
		log.Printf("Database migration failed: %v", err)
		return nil, err
	}
	log.Println("🚀 PostgreSQL connection and migration successful")
	return db, nil
}

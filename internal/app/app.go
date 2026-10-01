package app

import (
	"log"

	"github.com/mamun-jsx/Ticket-Booking-App/config"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/database"
	"gorm.io/gorm"
)

// App holds the core dependencies of the application like configuration and database connection.
type App struct {
	Config *config.Config
	DB     *gorm.DB
}

// BootStrapApp loads application configuration and establishes the database connection.
func BootStrapApp() *App {
	// 1. Load environment variables into configuration struct
	cfg := config.LoadEnv()

	// 2. Establish connection to the PostgreSQL database and run auto-migrations
	db, err := database.DatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
		return nil
	}

	// 3. Return the initialized App instance
	return &App{
		Config: cfg,
		DB:     db,
	}
}

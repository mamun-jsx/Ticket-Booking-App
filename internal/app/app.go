package app

import (
	"log"

	"github.com/mamun-jsx/Ticket-Booking-App/config"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/database"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Models *ModelsRec
}

func BootStrapApp() *App {
	cfg := config.LoadEnv()
	db, err := database.DatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
		return nil
	}

	appMod := initModels(db) // coming from ModelsReceiver and which holds all models and handlers

	return &App{
		Config: cfg,
		DB:     db,
		Models: appMod,
	}
}

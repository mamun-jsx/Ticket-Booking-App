package app

import (
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user"
	"gorm.io/gorm"
)

type ModelsRec struct {
	UserHandler *user.Handler
}

func initModels(db *gorm.DB) *ModelsRec {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	return &ModelsRec{
		UserHandler: userHandler,
	}
}

package app

import "gorm.io/gorm"

type ModelsRec struct {
}

func initModels(db *gorm.DB) *ModelsRec {
	return &ModelsRec{}
}

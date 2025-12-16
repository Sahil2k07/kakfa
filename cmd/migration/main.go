package main

import (
	"github.com/Sahil2k07/kakfa/internal/configs"
	"github.com/Sahil2k07/kakfa/internal/connections"
	"github.com/Sahil2k07/kakfa/internal/models"
	"github.com/charmbracelet/log"
)

func init() {
	configs.LoadConfig()
	connections.ConnectWDB()
}

func main() {
	// Migration - 1
	models := []any{&models.User{}, &models.Profile{}, &models.Todo{}}

	err := connections.WDB.AutoMigrate(models...)
	if err != nil {
		log.Errorf("Migration failed: %v", err)
		return
	}

	log.Infof("Migration Completed Successfully!")
}

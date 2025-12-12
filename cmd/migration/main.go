package main

import (
	"github.com/Sahil2k07/kakfa/internal/configs"
	"github.com/Sahil2k07/kakfa/internal/database"
	"github.com/charmbracelet/log"
)

func init() {
	configs.LoadConfig()
	database.ConnectWDB()
}

func main() {
	// Migration - 1
	models := []any{}

	err := database.WDB.AutoMigrate(models...)
	if err != nil {
		log.Errorf("Migration failed: %v", err)
		return
	}

	log.Infof("Migration Completed Successfully!")
}

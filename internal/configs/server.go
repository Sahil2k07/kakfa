package configs

import (
	"os"
	"strings"
)

type serverConfig struct {
	Port    string   `toml:"port"`
	Origins []string `toml:"origins"`
}

func loadServerConfig() serverConfig {
	org := os.Getenv("ORIGINS")
	return serverConfig{
		Port:    os.Getenv("PORT"),
		Origins: strings.Split(org, ","),
	}
}

func GetServerConfig() serverConfig {
	return serverConfig{
		Port:    globalConfig.Server.Port,
		Origins: globalConfig.Server.Origins,
	}
}

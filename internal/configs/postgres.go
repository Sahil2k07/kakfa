package configs

import (
	"fmt"
	"os"
)

type postgresConfig struct {
	Host     string `toml:"pg_host"`
	Port     string `toml:"pg_port"`
	User     string `toml:"pg_user"`
	Password string `toml:"pg_password"`
	Name     string `toml:"pg_name"`
}

func loadPostgresConfig() postgresConfig {
	return postgresConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Name:     os.Getenv("PG_NAME"),
	}
}

func getPostgresConfig() postgresConfig {
	return postgresConfig{
		Host:     globalConfig.Postgres.Host,
		Port:     globalConfig.Postgres.Port,
		User:     globalConfig.Postgres.User,
		Password: globalConfig.Postgres.Password,
		Name:     globalConfig.Postgres.Name,
	}
}

func GetPostgresConfig() string {
	conf := getPostgresConfig()

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
}

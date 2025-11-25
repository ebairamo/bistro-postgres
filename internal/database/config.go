package database

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
	PortApp  int
	DataDir  string
}

func LoadConfig() *Config {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	portStr := os.Getenv("DB_PORT")
	portAppStr := os.Getenv("PORT")
	dataDir := os.Getenv("DATA_DIR")
	port, _ := strconv.Atoi(portStr)
	portApp, _ := strconv.Atoi(portAppStr)

	cfg := &Config{
		Host:     host,
		User:     user,
		Password: password,
		Name:     name,
		Port:     port,
		PortApp:  portApp,
		DataDir:  dataDir,
	}

	slog.Info("Config loaded", "host", host, "db", name, "port", port)
	return cfg
}

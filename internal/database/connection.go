package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(cfg *Config) (*sql.DB, error) {

	conStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	conn, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil

}

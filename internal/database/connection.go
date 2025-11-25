package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(cfg *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	var db *sql.DB
	var err error

	// Пытаемся подключиться до 5 раз с задержкой
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				slog.Info("Connected to PostgreSQL")
				return db, nil
			}
		}

		slog.Warn("Connection attempt failed", "attempt", i+1, "error", err)
		time.Sleep(2 * time.Second) // ждём 2 секунды перед новой попыткой
	}

	return nil, err
}

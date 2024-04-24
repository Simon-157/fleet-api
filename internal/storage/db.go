package storage

import (
	"database/sql"
	"fleet_api/config"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

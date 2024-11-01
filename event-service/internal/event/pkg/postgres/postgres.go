package postgres 

import (
	"database/sql"
	"fmt"
	config "event-service/internal/event/pkg/load"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	cfg.Postgres.Host, 
	cfg.Postgres.Port, 
	cfg.Postgres.User, 
	cfg.Postgres.Password, 
	cfg.Postgres.Database)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
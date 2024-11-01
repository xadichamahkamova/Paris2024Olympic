package postgres

import (
	"database/sql"
	"fmt"
	config "user-service/internal/user/pkg/load"

	_ "github.com/lib/pq"
)

func InitDB(cfg config.Config) (*sql.DB, error) {
	target := fmt.Sprintf(
		`
			host=%s
			port=%d
			user=%s
			password=%s
			dbname=%s
			sslmode=disable	
		`,
		cfg.Postgres.Host, 
		cfg.Postgres.Port, 
		cfg.Postgres.User, 
		cfg.Postgres.Password, 
		cfg.Postgres.Database,
	)

	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package psql

import (
	_ "database/sql"
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"http-pattern/internal/config"
)

type Storage struct {
	db *sqlx.DB
}

func New(cfg config.Postgres) (*Storage, error) {

	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			cfg.User, cfg.Password, cfg.DatabaseName))
	return &Storage{db: db}, err
}

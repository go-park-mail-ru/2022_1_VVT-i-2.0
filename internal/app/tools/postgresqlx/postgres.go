package postgresqlx

import (
	"fmt"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresqlX(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d", cfg.User, cfg.DbName, cfg.Password, cfg.Host, cfg.Port)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.ConnectionMax)
	return db, nil
}

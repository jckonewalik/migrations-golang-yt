package database

import (
	"database/sql"
	"fmt"

	"github.com/jckonewalik/migration-demo-yt/config"
)

type postgresDB struct {
	Db *sql.DB
}

func NewPosgresDB() (*postgresDB, error) {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Env.DbUser, config.Env.DbPassword, config.Env.DbServer, config.Env.DbPort, config.Env.DbDatabase)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return &postgresDB{Db: db}, nil
}

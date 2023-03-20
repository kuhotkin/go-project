package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connection struct {
	*sql.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SSLMode  string
}

func NewPostgresDB(config Config) (*Connection, error) {
	str := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s", config.Host, config.Port, config.DbName, config.SSLMode)

	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Connection{
		DB: db,
	}, nil
}

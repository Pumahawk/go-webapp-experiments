package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConf struct {
	User     string
	Password string
	DBName   string
}

func CreateDatabaseConnection(dbConf DBConf) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.User, dbConf.Password, dbConf.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable initialize driver. %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to connect to database. %w", err)
	}

	return db, nil
}

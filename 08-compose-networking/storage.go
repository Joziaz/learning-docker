package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	port = 5432
)

var (
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
	dbName   = os.Getenv("POSTGRES_DB")
)

func newDbConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		errMsg := fmt.Errorf("failed to connect to database: %w", err)
		err = errors.Join(errMsg, err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		errMsg := fmt.Errorf("ping to database failed: %w", err)
		err = errors.Join(errMsg, err)
		return nil, err
	}

	return db, nil
}

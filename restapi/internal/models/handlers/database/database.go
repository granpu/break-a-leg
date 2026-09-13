package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sqlx.Connect()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdConns(5)

	return db, nil
}

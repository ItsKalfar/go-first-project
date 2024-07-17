package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func MySQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal("Db connection was not possible", err)
	}

	return db, nil
}
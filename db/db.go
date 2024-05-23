package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// create a new MySQL Storage, which receives a configuration cfg
// open a SQL db
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error){
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver

	"github.com/lbergamim-daitan/golang-rump-up/internal/config"
)

type Database struct {
	DB *sql.DB
}

type DatabaseInterface interface {
	ConnecterSQL() error
}

func (d *Database) ConnecterSQL() error {
	db, err := sql.Open("mysql", config.DatabaseStringConnection)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}
	d.DB = db
	return err
}

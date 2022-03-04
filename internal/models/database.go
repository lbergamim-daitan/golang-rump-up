package models

import (
	"github.com/lbergamim-daitan/golang-rump-up/internal/config"
	"github.com/lbergamim-daitan/golang-rump-up/internal/mysql"
	"github.com/lbergamim-daitan/golang-rump-up/internal/orm"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

type DatabaseInterface interface {
	Connect() error
	Insert(table string, columnName string, value string) (uint64, error)
	InsertMany(table string, columnName string, rows [][]string, ID uint64) error
	Query(table string, columnName string, value string) ([]responses.DefaultQuery, error)
	QueryAvailable(table string, columnName string, ID string) (responses.DefaultQuery, error)
	QueryCount(table string, columnName string) ([]responses.DefaultQuery, error)
	QueryID(table string, value string) ([]responses.DefaultQuery, error)
	Update(table string, columnName string, value string, ID string) (uint64, error)
	Delete(table string, ID string) error
}

func DatabaseChoose() DatabaseInterface {
	if config.DBImplem == "orm" {
		return &orm.ORM{}
	}
	return &mysql.Mysql{}
}

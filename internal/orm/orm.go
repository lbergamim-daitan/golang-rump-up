package orm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver

	"github.com/lbergamim-daitan/golang-rump-up/internal/config"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

type ORM struct {
	DB *sql.DB
}

func (d *ORM) Connect() error {
	db, err := sql.Open("mysql", config.ORMConnection)
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

func (d *ORM) Insert(table string, columnName string, value string) (uint64, error) {
	x := fmt.Sprintf("insert into %s (%s) values(?)", table, columnName)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return 0, err
	}
	defer d.DB.Close()

	result, err := statement.Exec(value)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (d *ORM) InsertMany(table string, columnName string, rows [][]string, ID uint64) error {
	x := fmt.Sprintf("insert into %s (%s) values(?, ?)", table, columnName)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return err
	}
	defer d.DB.Close()

	for _, row := range rows {
		_, err := statement.Exec(row[0], ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *ORM) Query(table string, columnName string, values string) ([]responses.DefaultQuery, error) {
	x := fmt.Sprintf("select * from %s where %s LIKE ?", table, columnName)
	query, err := d.DB.Query(x, values)
	if err != nil {
		return nil, err
	}
	defer d.DB.Close()

	var defaultQuerys []responses.DefaultQuery

	for query.Next() {
		var defaultQuery responses.DefaultQuery

		if err = query.Scan(
			&defaultQuery.ID,
			&defaultQuery.Name,
		); err != nil {
			return nil, err
		}

		defaultQuerys = append(defaultQuerys, defaultQuery)
	}

	return defaultQuerys, nil
}

func (d *ORM) QueryAvailable(table string, columnName string, ID string) (responses.DefaultQuery, error) {
	var defaultQuery responses.DefaultQuery

	x := fmt.Sprintf("select * from %s where %s = ? order by rand() limit 1", table, columnName)

	_, err := d.DB.Begin()
	if err != nil {
		return defaultQuery, err
	}

	row := d.DB.QueryRow(x, ID)

	if err := row.Scan(
		&defaultQuery.ID,
		&defaultQuery.CompanyID,
		&defaultQuery.Number,
	); err != nil {
		return defaultQuery, err
	}

	return defaultQuery, nil
}

func (d *ORM) QueryID(table string, value string) ([]responses.DefaultQuery, error) {
	x := fmt.Sprintf("select * from %s where ID = ?", table)
	query, err := d.DB.Query(x, value)
	if err != nil {
		return nil, err
	}
	defer d.DB.Close()

	var defaultQuerys []responses.DefaultQuery

	for query.Next() {
		var defaultQuery responses.DefaultQuery

		if err = query.Scan(
			&defaultQuery.ID,
			&defaultQuery.Name,
		); err != nil {
			return nil, err
		}

		defaultQuerys = append(defaultQuerys, defaultQuery)
	}

	return defaultQuerys, nil
}

func (d *ORM) QueryCount(table string, columnName string) ([]responses.DefaultQuery, error) {
	x := fmt.Sprintf("select %s, count(*) AS `available_phones` from %s group by %s", columnName, table, columnName)
	query, err := d.DB.Query(x)
	if err != nil {
		return nil, err
	}
	defer d.DB.Close()

	var defaultQuerys []responses.DefaultQuery

	for query.Next() {
		var defaultQuery responses.DefaultQuery

		if err = query.Scan(
			&defaultQuery.CompanyID,
			&defaultQuery.PhoneQuantity,
		); err != nil {
			return nil, err
		}

		defaultQuerys = append(defaultQuerys, defaultQuery)
	}

	return defaultQuerys, nil
}

func (d *ORM) Update(table string, columnName string, value string, ID string) (uint64, error) {
	x := fmt.Sprintf("update %s set %s = ? where ID = ?", table, columnName)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return 0, err
	}
	defer d.DB.Close()

	result, err := statement.Exec(value, ID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (d *ORM) Delete(table string, ID string) error {
	x := fmt.Sprintf("delete from %s where ID = ?", table)
	statement, err := d.DB.Prepare(x)

	if err != nil {
		return err
	}
	defer d.DB.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

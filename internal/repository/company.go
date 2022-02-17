package repository

import (
	"fmt"

	"github.com/lbergamim-daitan/golang-rump-up/internal/db"
	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
)

type Company struct {
	database *db.Database
}

func NewCompanyRepo(database *db.Database) models.CompanyInterface {
	return &Company{database}
}

func (c *Company) Create(company models.Company) (uint64, error) {
	err := c.database.ConnecterSQL()
	if err != nil {
		return 0, err
	}

	statement, err := c.database.DB.Prepare("insert into company (name) values(?)")
	if err != nil {
		return 0, err
	}
	defer c.database.DB.Close()

	result, err := statement.Exec(company.Name)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (c *Company) List(name string) ([]models.Company, error) {
	err := c.database.ConnecterSQL()
	if err != nil {
		return nil, err
	}
	name = fmt.Sprintf("%%%s%%", name)

	query, err := c.database.DB.Query(
		"select * from company where name LIKE ?", name)

	if err != nil {
		return nil, err
	}
	defer c.database.DB.Close()

	var companies []models.Company

	for query.Next() {
		var company models.Company

		if err = query.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (c *Company) ListID(ID string) ([]models.Company, error) {
	err := c.database.ConnecterSQL()
	if err != nil {
		return nil, err
	}
	query, err := c.database.DB.Query(
		"select * from company where ID = ?", ID)

	if err != nil {
		return nil, err
	}
	defer c.database.DB.Close()

	var companies []models.Company

	for query.Next() {
		var company models.Company

		if err = query.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (c *Company) Update(ID string, company models.Company) (uint64, error) {
	err := c.database.ConnecterSQL()
	if err != nil {
		return 0, err
	}
	statement, err := c.database.DB.Prepare("update company set name = ? where ID = ?")
	if err != nil {
		return 0, err
	}
	defer c.database.DB.Close()

	result, err := statement.Exec(company.Name, ID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (c *Company) Delete(ID string) error {
	err := c.database.ConnecterSQL()
	if err != nil {
		return err
	}
	statement, err := c.database.DB.Prepare(
		"delete from company where ID = ?")

	if err != nil {
		return err
	}
	defer c.database.DB.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

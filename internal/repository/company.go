package repository

import (
	"fmt"

	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

type Company struct {
	database models.DatabaseInterface
}

func NewCompanyRepo(database models.DatabaseInterface) models.CompanyInterface {
	return &Company{database}
}

func (c *Company) Create(company models.Company) (uint64, error) {
	err := c.database.Connect()
	if err != nil {
		return 0, err
	}

	lastID, err := c.database.Insert("company", "name", company.Name)
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (c *Company) List(name string) ([]responses.DefaultQuery, error) {
	err := c.database.Connect()
	if err != nil {
		return nil, err
	}
	name = fmt.Sprintf("%%%s%%", name)

	query, err := c.database.Query("company", "name", name)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (c *Company) ListID(ID string) ([]responses.DefaultQuery, error) {
	err := c.database.Connect()
	if err != nil {
		return nil, err
	}
	query, err := c.database.QueryID("company", ID)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (c *Company) Update(ID string, company models.Company) (uint64, error) {
	err := c.database.Connect()
	if err != nil {
		return 0, err
	}

	lastID, err := c.database.Update("company", "name", company.Name, ID)
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (c *Company) Delete(ID string) error {
	err := c.database.Connect()
	if err != nil {
		return err
	}

	err = c.database.Delete("company", ID)
	if err != nil {
		return err
	}

	return nil
}

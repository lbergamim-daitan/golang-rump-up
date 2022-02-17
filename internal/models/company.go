package models

import (
	"errors"
	"strings"
)

type CompanyInterface interface {
	Create(company Company) (uint64, error)
	List(name string) ([]Company, error)
	ListID(ID string) ([]Company, error)
	Update(ID string, company Company) (uint64, error)
	Delete(ID string) error
}

type Company struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//Prepare call methods to validate company
func (c *Company) Prepare() error {
	if err := c.validate(); err != nil {
		return err
	}

	c.format()
	return nil
}

func (c *Company) validate() error {
	if c.Name == "" {
		return errors.New("name is mandatory")
	}

	return nil
}

func (c *Company) format() {
	c.Name = strings.TrimSpace(c.Name)
}

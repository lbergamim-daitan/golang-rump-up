package models

import (
	"errors"
	"strings"

	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

type PhonesInterface interface {
	Create(phone Phone) error
	ListAvailable(ID string) (responses.DefaultQuery, error)
	List() ([]responses.DefaultQuery, error)
}

type Phone struct {
	ID            uint64     `json:"id,omitempty"`
	CompanyID     uint64     `json:"company_id,omitempty"`
	Number        string     `json:"phone_number,omitempty"`
	Rows          [][]string `json:"rows,omitempty"`
	PhoneQuantity uint64     `json:"available_phones,omitempty"`
}

func (p *Phone) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *Phone) ValidateFile() error {
	if len(p.Rows) > 5000000 {
		return errors.New("exceeded file size")
	}

	return nil
}

func (p *Phone) validate() error {
	if p.Number == "" {
		return errors.New("name is mandatory")
	}

	return nil
}

func (p *Phone) format() {
	p.Number = strings.TrimSpace(p.Number)
}

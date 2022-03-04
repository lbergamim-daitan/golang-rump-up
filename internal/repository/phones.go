package repository

import (
	"strconv"

	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

type Phone struct {
	database models.DatabaseInterface
}

func NewPhoneRepo(database models.DatabaseInterface) models.PhonesInterface {
	return &Phone{database}
}

func (p *Phone) ListAvailable(ID string) (responses.DefaultQuery, error) {
	err := p.database.Connect()
	if err != nil {
		return responses.DefaultQuery{}, err
	}

	phone, err := p.database.QueryAvailable("phone", "company_id", ID)
	if err != nil {
		return responses.DefaultQuery{}, err
	}

	err = p.database.Delete("phone", strconv.FormatUint(phone.ID, 10))
	if err != nil {
		return responses.DefaultQuery{}, err
	}

	return phone, nil
}

func (p *Phone) List() ([]responses.DefaultQuery, error) {
	err := p.database.Connect()
	if err != nil {
		return []responses.DefaultQuery{}, err
	}

	phones, err := p.database.QueryCount("phone", "company_id")
	if err != nil {
		return []responses.DefaultQuery{}, err
	}

	return phones, nil
}

func (p *Phone) Create(phone models.Phone) error {
	err := p.database.Connect()
	if err != nil {
		return err
	}

	err = p.database.InsertMany("phone", "number, company_id", phone.Rows, phone.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

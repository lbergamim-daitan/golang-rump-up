package repository

import (
	"errors"

	"github.com/lbergamim-daitan/golang-rump-up/internal/db"
	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
)

type Phone struct {
	database *db.Database
}

func NewPhoneRepo(database *db.Database) models.PhonesInterface {
	return &Phone{database}
}

func (p *Phone) List() ([]models.Phone, error) {
	err := p.database.ConnecterSQL()
	if err != nil {
		return nil, err
	}

	query, err := p.database.DB.Query(
		"select company_id, count(*) AS `available_phones` from phone group by company_id;")

	if err != nil {
		return nil, err
	}
	defer p.database.DB.Close()

	var phones []models.Phone

	for query.Next() {
		var phone models.Phone

		if err = query.Scan(
			&phone.CompanyID,
			&phone.PhoneQuantity,
		); err != nil {
			return nil, err
		}

		phones = append(phones, phone)
	}

	return phones, nil
}

func (p *Phone) ListAvailable(ID string) (models.Phone, error) {
	var phone models.Phone
	err := p.database.ConnecterSQL()
	if err != nil {
		return phone, err
	}

	transaction, err := p.database.DB.Begin()
	if err != nil {
		return phone, err
	}

	row := p.database.DB.QueryRow(
		"select * from phone where company_id = ? order by rand() limit 1", ID)

	if err := row.Scan(
		&phone.ID,
		&phone.CompanyID,
		&phone.Number,
	); err != nil {
		return models.Phone{}, err
	}

	err = p.Delete(phone.ID)
	if err != nil {
		return models.Phone{}, err
	}

	err = transaction.Rollback()
	if err != nil {
		return models.Phone{}, err
	}

	return phone, nil
}

func (p *Phone) Delete(ID uint64) error {
	deleteQuery, err := p.database.DB.Exec("delete from phone where id = ?;", ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := deleteQuery.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (p *Phone) Create(phone models.Phone) error {
	err := p.database.ConnecterSQL()
	if err != nil {
		return err
	}

	statement, err := p.database.DB.Prepare("insert into phone (number, company_id) values(?, ?)")
	if err != nil {
		return err
	}
	defer p.database.DB.Close()

	for _, row := range phone.Rows {
		_, err := statement.Exec(row[0], phone.CompanyID)
		if err != nil {
			return err
		}
	}

	return nil
}

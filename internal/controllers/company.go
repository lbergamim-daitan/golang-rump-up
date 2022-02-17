package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lbergamim-daitan/golang-rump-up/internal/db"
	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
	"github.com/lbergamim-daitan/golang-rump-up/internal/repository"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var company models.Company
	if err = json.Unmarshal(bodyRequest, &company); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := company.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	database := db.Database{}
	companyRepository := repository.NewCompanyRepo(&database)
	company.ID, err = companyRepository.Create(company)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, company)
}

func ListCompanies(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	database := db.Database{}
	companyRepository := repository.NewCompanyRepo(&database)

	companies, err := companyRepository.List(name)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, companies)
}

func ListCompany(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	database := db.Database{}
	companyRepository := repository.NewCompanyRepo(&database)

	companies, err := companyRepository.ListID(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	httpStatusCode := http.StatusOK
	if len(companies) == 0 {
		httpStatusCode = http.StatusNotFound
		responses.JSON(w, httpStatusCode, "resource not found")
		return
	}

	responses.JSON(w, httpStatusCode, companies)

}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var company models.Company
	if err = json.Unmarshal(bodyRequest, &company); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := company.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	database := db.Database{}
	companyRepository := repository.NewCompanyRepo(&database)
	company.ID, err = companyRepository.Update(ID, company)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, company)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	database := db.Database{}
	companyRepository := repository.NewCompanyRepo(&database)

	err := companyRepository.Delete(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, "")

}

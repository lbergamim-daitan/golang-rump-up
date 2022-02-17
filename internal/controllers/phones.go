package controllers

import (
	"encoding/csv"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lbergamim-daitan/golang-rump-up/internal/db"
	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
	"github.com/lbergamim-daitan/golang-rump-up/internal/repository"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

func AvailablePhones(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	database := db.Database{}
	phoneRepository := repository.NewPhoneRepo(&database)

	phone, err := phoneRepository.ListAvailable(ID)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	httpStatusCode := http.StatusOK
	if reflect.DeepEqual(phone, models.Phone{}) {
		httpStatusCode = http.StatusNotFound
		responses.JSON(w, httpStatusCode, "error to process request")
		return
	}

	responses.JSON(w, httpStatusCode, phone)
}

func ListPhones(w http.ResponseWriter, r *http.Request) {

	database := db.Database{}
	companyRepository := repository.NewPhoneRepo(&database)

	phones, err := companyRepository.List()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, phones)
}

func CreatePhones(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	var phone models.Phone

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer file.Close()

	reads := csv.NewReader(file)
	phone.Rows, err = reads.ReadAll()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	phone.CompanyID, err = strconv.ParseUint(ID, 10, 64)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = phone.ValidateFile()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	database := db.Database{}
	phoneRepository := repository.NewPhoneRepo(&database)

	err = phoneRepository.Create(phone)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, "file uploaded")
}

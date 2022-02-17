package repository

import (
	"reflect"
	"testing"

	"github.com/lbergamim-daitan/golang-rump-up/internal/db"
)

func TestNewCompanyRepo(t *testing.T) {
	database := db.Database{}

	repository := NewCompanyRepo(&database)

	if reflect.TypeOf(repository) != reflect.TypeOf(&Company{&database}) {
		t.Error()
	}

}

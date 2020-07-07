package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/medson/superheroesAPI/src/application"
	"github.com/medson/superheroesAPI/src/application/config"
	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/medson/superheroesAPI/src/infra/repository"
)

func TestListSupers(t *testing.T) {
	var newSuper models.Super
	var result []models.Super

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	newSuper.Name = "Gohan"
	newSuper.FullName = "Son Gohan"
	repository.SuperRepository{}.Create(&newSuper)

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)

	expected := newSuper.Name

	if len(result[0].Name) <= 0 {
		t.Errorf(defaultError, expected, result[0].Name)
	}
}

func TestListSuperByName(t *testing.T) {
	var result []models.Super
	var newSuper models.Super

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	newSuper.Name = "Dama de Ferro"
	newSuper.FullName = "Pepper Potts"
	repository.SuperRepository{}.Create(&newSuper)

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("name", "Dama de Ferro")
	req.URL.RawQuery = queryParams.Encode()

	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)

	expected := newSuper.Name

	if result[0].Name != newSuper.Name {
		t.Errorf(defaultError, expected, result[0].Name)
	}
}

func TestListSuperByNameNotRegistered(t *testing.T) {
	var obj models.ErrorResponse

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("name", "Vegeto")
	req.URL.RawQuery = queryParams.Encode()
	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &obj)
	expected := "Super not found"

	if obj.Error != expected {
		t.Errorf(defaultError, expected, obj.Error)
	}
}

func TestListSuperByIDNotRegistered(t *testing.T) {
	var result models.ErrorResponse

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("id", "00000000-0000-0000-0000-000000000000")
	req.URL.RawQuery = queryParams.Encode()

	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)

	expected := "Super not found"

	if result.Error != expected {
		t.Errorf(defaultError, expected, result.Error)
	}
}

func TestListSuperBySideInvalid(t *testing.T) {
	var result models.ErrorResponse

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("superSide", "empty")
	req.URL.RawQuery = queryParams.Encode()

	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)

	expected := "Invalid Param superSide, need to be bad or good."

	if result.Error != expected {
		t.Errorf(defaultError, expected, result.Error)
	}
}

func TestListSuperByID(t *testing.T) {
	var newSuper models.Super
	var result []models.Super

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	newSuper.Name = "Majin Boo"
	newSuper.FullName = "Majin Boo"
	repository.SuperRepository{}.Create(&newSuper)

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("name", "Majin Boo")
	req.URL.RawQuery = queryParams.Encode()

	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)
	expected := "Majin Boo"

	if result[0].Name != newSuper.Name {
		t.Errorf(defaultError, expected, result[0].Name)
	}
}

func TestListSupersBySide(t *testing.T) {
	var newSuper models.Super
	var result []models.Super

	app := application.NewServer()
	config.Settings.Enviroment = "test"

	newSuper.Name = "Freeza"
	newSuper.FullName = "Lord Freeza"
	newSuper.Alignment = "bad"
	repository.SuperRepository{}.Create(&newSuper)

	req, _ := http.NewRequest("GET", "/api/v1/supers", nil)
	queryParams := req.URL.Query()
	queryParams.Add("superSide", "bad")
	req.URL.RawQuery = queryParams.Encode()

	res, _ := app.Test(req, -1)
	bodyBytes, _ := (ioutil.ReadAll(res.Body))
	json.Unmarshal(bodyBytes, &result)
	expected := "bad"

	if (result[0].Name != newSuper.Name) && (result[0].Alignment != "bad") {
		t.Errorf(defaultError, expected, result[0].Alignment)
	}
}

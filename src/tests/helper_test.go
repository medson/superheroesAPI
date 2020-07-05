package models

import (
	"testing"

	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/medson/superheroesAPI/src/helpers"
)

const defaultError = "Expected value %v, but the result found was %v."

func TestIfStringSplittingIsCorrect(t *testing.T) {
	arrange := helpers.SplitByOutsideCommas("Casey and Elaine Griggs (foster parents, deceased),\nLena Luthor (sister pre-Crisis; daughter post-Crisis),\nElizabeth Perske (ex-wife),\nPerry J. White Jr. (son, deceased),\nContessa Erica Alexandra del Portenza (wife, assumed deceased)")

	result := len(arrange)
	expected := 5

	if result != expected {
		t.Errorf(defaultError, expected, result)
	}
}

func TestCorrectModelsMapping(t *testing.T) {
	superResponse := models.SuperResponse{Name: "Goku"}
	var super models.Super

	arrange := helpers.MapResponseToModel(superResponse, super, make([]string, 0))

	result := arrange.Name
	expected := "Goku"

	if result != expected {
		t.Errorf(defaultError, expected, result)
	}
}

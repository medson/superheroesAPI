package external

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/medson/superheroesAPI/src/domain/models"
)

// CallSuperAPI is responsible for making requests to the SuperHeroesAPI
//
// superName - Field that represent the super hero name to be search
//
// Returns a SuperAPIResponse model with data obtained by request
//
func CallSuperAPI(superName string) (models.SuperAPIResponse, models.ErrorResponse) {
	var responseData models.SuperAPIResponse
	err := models.ErrorResponse{}
	res, fail := http.Get("https://superheroapi.com/api/3256918617734157/search/" + superName)

	if fail != nil {
		log.Fatal(fail)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&responseData)

	if responseData.Response == "error" {
		err.Error = responseData.Error
	}

	return responseData, err
}

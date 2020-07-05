package helpers

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/valyala/fasthttp"
)

// GetQueryParams is responsible to format the query params selected
// in request.
//
// queryArgs - Field that represent query arguments
//
// Returns a map with filter key and value
//
func GetQueryParams(queryArgs *fasthttp.Args) map[string]string {
	filter := make(map[string]string)

	filter["superSide"] = ""
	filter["id"] = ""
	filter["name"] = ""

	if queryArgs.Has("superSide") {
		filter["superSide"] = string(queryArgs.Peek("superSide"))
	} else if queryArgs.Has("id") {
		filter["id"] = string(queryArgs.Peek("id"))
	} else if queryArgs.Has("name") {
		filter["name"] = string(queryArgs.Peek("name"))
	}

	return filter
}

// SplitByOutsideCommas receive a string and separate in a string list, finding
// by patterns like "," and "),".
//
// stringToParse - Field that represent string that will be splited
//
// Returns a list with strings
//
func SplitByOutsideCommas(stringToParse string) []string {
	var cleanedSlice []string
	parsedSlice := regexp.MustCompile("\\(.*?\\)|(,)").Split(stringToParse, -1)

	for currentSuperRelative := 0; currentSuperRelative < len(parsedSlice); currentSuperRelative++ {
		if parsedSlice[currentSuperRelative] != "" {
			cleanedSlice = append(cleanedSlice, parsedSlice[currentSuperRelative])
		}
	}
	return cleanedSlice
}

// ShowDocs its responsible to serve the swagger.json to be used by go-swagger package.
//
// c - A pointer with fiber context
//
// Returns a swagger file deserialized.
//
func ShowDocs(c *fiber.Ctx) {
	raw, err := ioutil.ReadFile("./docs/swagger.json")
	var result map[string]interface{}
	json.Unmarshal([]byte(raw), &result)
	if err != nil {
		c.Send("Error")
	}
	c.JSON(result)
}

// MapResponseToModel its responsible to map response obtained by external
//request to entity that will be stored at database.
//
// superResponse - Represents the external response
// super         - Represents the super that will store mapped data
// relatives     - List of relatives mapped
//
// Returns a super mapped ready to be inserted at database
//
func MapResponseToModel(superResponse models.SuperResponse, super models.Super, relatives []string) models.Super {
	super.Name = superResponse.Name
	super.FullName = superResponse.Biography.FullName
	super.Intelligence, _ = strconv.Atoi(superResponse.Powerstats.Intelligence)
	super.Power, _ = strconv.Atoi(superResponse.Powerstats.Power)
	super.Occupation = superResponse.Work.Occupation
	super.Image = superResponse.Image.URL
	super.Alignment = superResponse.Biography.Alignment
	super.RelativesAmount = len(relatives)
	super.Relatives = superResponse.Connections.Relatives
	return super
}

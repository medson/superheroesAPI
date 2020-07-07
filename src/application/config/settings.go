package config

import (
	"os"

	"github.com/victorgama/gopts"
)

// SettingsStruct represents the structure of enviroment variables
//
// APIPort     - Value of the port addres that the application will run
// DatabaseURL - String with db url connection
// SuperAPIURL - String with SuperHeroes external API URL
// DialectDB   - Db dialect that will be used by gorm
// Enviroment  - Current enviroment. Ex: dev,prod test.
//
type SettingsStruct struct {
	APIPort         string `default:"3000"`
	DatabaseURL     string `default:"postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"`
	SuperAPIURL     string `default:"https://superheroapi.com/api/3256918617734157/search/"`
	TestDatabaseURL string `default:"/tmp/gorm.db"`
	DialectDB       string `default:"postgres"`
	Enviroment      string `default:"dev"`
}

// Settings pointer with variables values
var Settings *SettingsStruct

func init() {
	set := gopts.LoadEnvs(SettingsStruct{}).(SettingsStruct)
	env := os.Getenv("ENV")

	if env == "test" {
		set.Enviroment = env
		set.DialectDB = "sqlite3"
	}

	Settings = &set
}

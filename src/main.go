package main

import (
	"github.com/medson/superheroesAPI/src/application"
	"github.com/medson/superheroesAPI/src/application/config"
)

// @title Super-Heroes API
// @version 1.0
// @description This is an API in Golang that integrate with https://superheroapi.com/
// @termsOfService http://swagger.io/terms/
// @contact.name Medson Mendes
// @contact.email medson095@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := application.NewServer()
	app.Listen(config.Settings.APIPort)
}

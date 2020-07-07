package infra

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lib/pq"
	"github.com/medson/superheroesAPI/src/application/config"
	"github.com/medson/superheroesAPI/src/domain/models"
)

// InitDB is responsible for initializing the connection to the database.
func InitDB() error {

	dbURL, DialectDB := setupDB()
	var err error

	if DialectDB == "postgres" {
		dbURL, err = pq.ParseURL(dbURL)
		if err != nil {
			log.Fatal(err)
		}
	}

	conn, err := gorm.Open(DialectDB, dbURL)

	if err != nil {
		return err
	}

	conn.AutoMigrate(&models.Super{}, &models.Group{})

	log.Println("Database Migrated")

	Database = conn
	return nil
}

// Database represents the configured database connection. This pointer will
// be available after `InitDB` is executed successfully
var Database *gorm.DB

// setupDB its responsible to setup database for respective enviroments
//
// Returns two strings with databaseURL and database dialect that will be used, ex.: "postgres".
func setupDB() (string, string) {
	if config.Settings.Enviroment == "test" {
		return config.Settings.TestDatabaseURL, config.Settings.DialectDB
	} else {
		return config.Settings.DatabaseURL, config.Settings.DialectDB
	}
}

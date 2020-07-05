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
	connStr, err := pq.ParseURL(config.Settings.DatabaseURL)
	if err != nil {
		return err
	}

	conn, err := gorm.Open("postgres", connStr)

	// conn, err := gorm.Open("sqlite3", ":memory:")
	// conn.Exec("PRAGMA foreign_keys = ON")

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

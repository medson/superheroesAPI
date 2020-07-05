package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base represents the entity on which other entities will be based
//
// ID 		 - A uuid field with the base entity unique identifier in database
// CreatedAt - A field that represent the instance creation time
// UpdatedAt - A field that represent the instance update time
// DeletedAt - A field that represent the instance delete time
//
type Base struct {
	ID        string     `sql:"type:uuid;primary_key`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate hook
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid.String())
}

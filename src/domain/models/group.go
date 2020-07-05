package models

// Group entity represents the groups list whose super belongs
//
// Base   - Field that represent the Base entity
// Name   - A field that represent the Group name
// Supers - A field that represent the supers related to that list
//
type Group struct {
	Base
	Name   string   `gorm:"size:255;not null;unique" json:"name"`
	Supers []*Super `gorm:"many2many:group_super;" json:"-"`
}

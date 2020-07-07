package repository

import (
	"os"

	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/medson/superheroesAPI/src/infra"
)

// SuperRepository its responsible to map response obtained by external
//request to entity that will be stored at database.
//
// superResponse - Represents the external response
// super         - Represents the super that will store mapped data
// relatives     - List of relatives mapped
//
// Returns a super mapped ready to be inserted at database
//
type SuperRepository struct {
}

// NewSuperRepository its responsible to instantiate SuperRepository
//request to entity that will be stored at database.
//
// superResponse - Represents the external response
// super         - Represents the super that will store mapped data
// relatives     - List of relatives mapped
//
// Returns a super mapped ready to be inserted at database
//
func NewSuperRepository() SuperRepository {
	return SuperRepository{}
}

// FindAll list all supers at database
//
// Returns a list with supers stored at database
//
func (superRepo SuperRepository) FindAll() []models.Super {
	var supers []models.Super
	infra.Database.Preload("GroupsAssociated").Find(&supers)
	return supers
}

// Create insert one super at database
//
// super - A super to be stored at database
//
// Returns a super stored at database
//
func (superRepo SuperRepository) Create(super *models.Super) models.Super {
	var newSuper models.Super
	infra.Database.Create(&super).First(&newSuper)
	return newSuper
}

// FindByID search one super by id
//
// superID - A super hero or villain ID to search at database
//
// Returns a super stored at database
//
func (superRepo SuperRepository) FindByID(superID string) models.Super {
	var newSuper models.Super
	infra.Database.Preload("GroupsAssociated").Where("id =?", superID).Find(&newSuper)
	return newSuper
}

// DeleteOne delete one super by id
//
// super - A instance at super in database
//
func (superRepo SuperRepository) DeleteOne(super models.Super) {
	infra.Database.Delete(&super)
}

// FindByName search one super by name
//
//  - A super hero or villain name to search at database
//
// Returns a super stored at database
//
func (superRepo SuperRepository) FindByName(superName string) models.Super {
	var super models.Super

	if os.Getenv("ENV") != "test" {
		infra.Database.Preload("GroupsAssociated").Where("name ILIKE ?", superName).First(&super)
	} else {
		infra.Database.Preload("GroupsAssociated").Where("name = ?", superName).First(&super)
	}

	return super
}

// FindByAlignment search one super by alignment
//
// alignment - A super hero or villain name to alignment at database
//
// Returns supers stored at database
//
func (superRepo SuperRepository) FindByAlignment(alignment string) []models.Super {
	var supers []models.Super
	infra.Database.Preload("GroupsAssociated").Where("alignment = ?", alignment).Find(&supers)
	return supers
}

// FindByNameAndFullName search one super by name and full name
//
// name - A super hero or villain name
// fullName - A super hero or villain full name
//
// Returns a super stored at database
//
func (superRepo SuperRepository) FindByNameAndFullName(fullName string, name string) models.Super {
	var super models.Super
	infra.Database.Where("full_name = ? AND name = ?", fullName, name).Find(&super)
	return super
}

// AddGroupsToSuper add a relation between supers and groups
//
// super - A super hero or villain that will receive a relation
// group - A group that will be related with a super
//
// Returns a super stored at database
//
func (superRepo SuperRepository) AddGroupsToSuper(super *models.Super, group models.Group) **models.Super {
	infra.Database.Where("name = ?", group.Name).FirstOrCreate(&group)

	infra.Database.Model(&super).Association("GroupsAssociated").Append(&group)
	return &super
}

package services

import (
	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/medson/superheroesAPI/src/helpers"
	"github.com/medson/superheroesAPI/src/infra/repository"
)

// SuperService represents a service layer that will call supers repository
//
// superRepo - Field that represent the supers repository layer
//
type SuperService struct {
	superRepo repository.SuperRepository
}

// Create insert one or more supers at database using repository
//
// SuperAPIResponse - Field that represent the supers found after call external service
//
// Returns a list with supers or super registered and a error, if happen.
//
func (s SuperService) Create(responseData models.SuperAPIResponse) ([]*models.Super, models.ErrorResponse) {
	supers := []*models.Super{}
	err := models.ErrorResponse{}

	for i := 0; i < len(responseData.Results); i++ {
		var super models.Super
		super = s.superRepo.FindByNameAndFullName(responseData.Results[i].Biography.FullName, responseData.Results[i].Name)

		if super.ID != "" {
			err.Error = "super already exists"
			return supers, err
		}

		relatives := helpers.SplitByOutsideCommas(responseData.Results[i].Connections.Relatives)
		currentGroups := helpers.SplitByOutsideCommas(responseData.Results[i].Connections.GroupAffiliation)
		super = helpers.MapResponseToModel(responseData.Results[i], super, relatives)

		s.superRepo.Create(&super)

		for _, groupName := range currentGroups {
			var newGroup models.Group
			newGroup.Name = groupName

			s.superRepo.AddGroupsToSuper(&super, newGroup)
		}

		supers = append(supers, &super)
	}
	return supers, err
}

// Get get one or more supers at database using repository
//
// filter - Field that represent query params that will be used for kind of search
//
// Returns a list with supers or super registered and a error, if happen.
//
func (s SuperService) Get(filter map[string]string) ([]models.Super, models.ErrorResponse) {
	var supers []models.Super
	var err models.ErrorResponse
	switch {
	case filter["superSide"] != "":
		if (filter["superSide"] != "good") && (filter["superSide"] != "bad") {
			err.Error = "Invalid Param superSide, need to be bad or good."
		} //||
		supers = s.superRepo.FindByAlignment(filter["superSide"])
	case filter["id"] != "":
		res := s.superRepo.FindByID(filter["id"])
		supers = append(supers, res)
	case filter["name"] != "":
		res := s.superRepo.FindByName(filter["name"])
		supers = append(supers, res)
	default:
		supers = s.superRepo.FindAll()
	}

	if (len(supers) > 0) && (supers[0].Name == "") {
		err.Error = "Super not found"
		return supers, err
	}
	return supers, err
}

// Destroy remove one or more supers stored at database
//
// superID - The identifier of super to be deleted
//
// Returns a message confirming that super has been deleted and a error, if happen.
//
func (s SuperService) Destroy(superID string) (string, models.ErrorResponse) {
	var super models.Super
	var err models.ErrorResponse

	super = s.superRepo.FindByID(superID)

	if super.Name == "" {
		err.Error = "Super not found"
		return "", err
	}

	s.superRepo.DeleteOne(super)
	return "Super deleted", err
}

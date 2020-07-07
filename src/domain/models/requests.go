package models

// SuperCreationReq represents the model of creation super request
//
// Name  - Field that represent the name of super to find
//
type SuperCreationReq struct {
	Name string `json:"name" validate:"required,string"`
}

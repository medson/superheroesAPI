package models

type SuperCreationReq struct {
	Name string `json:"name" validate:"required,string"`
}

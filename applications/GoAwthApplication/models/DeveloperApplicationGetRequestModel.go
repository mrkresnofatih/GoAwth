package models

type DeveloperApplicationGetRequestModel struct {
	DeveloperName          string `json:"developerName" validate:"required,min=6,max=30"`
	DeveloperApplicationId string `json:"developerApplicationId" validate:"required"`
}

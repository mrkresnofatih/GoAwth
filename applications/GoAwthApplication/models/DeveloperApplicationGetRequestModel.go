package models

type DeveloperApplicationReadRequestModel struct {
	DeveloperApplicationId string `json:"developerApplicationId" validate:"required"`
}

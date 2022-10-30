package models

type OauthRejectConsentRequestModel struct {
	DeveloperApplicationId string `json:"developerApplicationId" validate:"required"`
}

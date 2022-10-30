package models

type DeveloperApplicationCreateRequestModel struct {
	DeveloperName      string `json:"developerName" validate:"required,min=6,max=30"`
	Name               string `json:"name" validate:"required,min=6,max=50"`
	LogoUrl            string `json:"logoUrl"`
	SuccessRedirectUri string `json:"successRedirectUri" validate:"required,max=300"`
	FailedRedirectUri  string `json:"failedRedirectUri" validate:"required,max=300"`
}

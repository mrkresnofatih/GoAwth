package models

type DeveloperApplicationReadResponseModel struct {
	DeveloperApplicationId string `json:"developerApplicationId"`
	DeveloperName          string `json:"developerName"`
	Name                   string `json:"name"`
	LogoUrl                string `json:"logoUrl"`
	SuccessRedirectUri     string `json:"successRedirectUri"`
	FailedRedirectUri      string `json:"failedRedirectUri"`
	Secret                 string `json:"secret"`
}

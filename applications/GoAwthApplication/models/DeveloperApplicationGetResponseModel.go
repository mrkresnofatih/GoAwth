package models

type DeveloperApplicationGetResponseModel struct {
	DeveloperApplicationId string `json:"developerApplicationId"`
	DeveloperName          string `json:"developerName"`
	Name                   string `json:"name"`
	Secret                 string `json:"secret"`
	LogoUrl                string `json:"logoUrl"`
	SuccessRedirectUri     string `json:"successRedirectUri"`
	FailedRedirectUri      string `json:"failedRedirectUri"`
}

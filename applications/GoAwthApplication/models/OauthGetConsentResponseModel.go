package models

type OauthGetConsentResponseModel struct {
	DeveloperApplicationId       string            `json:"developerApplicationId"`
	DeveloperApplicationName     string            `json:"developerApplicationName"`
	DeveloperApplicationImageUrl string            `json:"developerApplicationImageUrl"`
	DeveloperName                string            `json:"developerName"`
	PlayerUsername               string            `json:"playerUsername"`
	PlayerImageUrl               string            `json:"playerImageUrl"`
	ScopeDefinitions             map[string]string `json:"scopeDefinitions"`
}

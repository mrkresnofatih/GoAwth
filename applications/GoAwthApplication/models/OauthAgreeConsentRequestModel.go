package models

type OauthAgreeConsentRequestModel struct {
	DeveloperApplicationId string `json:"developerApplicationId"`
	Scope                  string `json:"scope"`
	GrantType              string `json:"grantType"`
	PlayerUsername         string `json:"playerUsername"`
}

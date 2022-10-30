package models

type OauthAgreeConsentResponseModel struct {
	GrantId     string `json:"grantId"`
	RedirectUri string `json:"redirectUri"`
}

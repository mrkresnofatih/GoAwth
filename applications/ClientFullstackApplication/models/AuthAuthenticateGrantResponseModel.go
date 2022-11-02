package models

type AuthAuthenticateGrantResponseModel struct {
	GrantId         string `json:"grantId"`
	GrantToken      string `json:"grantToken"`
	PermittedScopes string `json:"permittedScopes"`
	ExpiresAt       string `json:"expiresAt"`
}

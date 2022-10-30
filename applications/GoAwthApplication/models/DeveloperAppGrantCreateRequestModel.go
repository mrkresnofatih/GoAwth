package models

type DeveloperAppGrantCreateRequestModel struct {
	PlayerUsername string
	ApplicationId  string
	ExpiresAt      string
	Scope          string
}

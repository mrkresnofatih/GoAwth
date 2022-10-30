package models

type DeveloperApplicationListRequestModel struct {
	DeveloperName string `json:"developerName" validate:"required,min=6,max=30"`
	Page          int    `json:"page" validate:"required,gt=0"`
	PageSize      int    `json:"pageSize" validate:"required,gt=0,lt=20"`
}

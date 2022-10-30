package models

type DeveloperApplicationListResponseModel struct {
	DeveloperName string                                 `json:"developerName"`
	Applications  []DeveloperApplicationGetResponseModel `json:"applications"`
}

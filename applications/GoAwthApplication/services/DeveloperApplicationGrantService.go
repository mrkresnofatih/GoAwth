package services

import (
	"github.com/google/uuid"
	"github.com/mrkresnofatih/go-awth/entities"
	"github.com/mrkresnofatih/go-awth/models"
	"gorm.io/gorm"
	"log"
)

type IDeveloperApplicationGrantService interface {
	Create(developerAppGrantCreateReq models.DeveloperAppGrantCreateRequestModel) (models.DeveloperAppGrantCreateResponseModel, error)
	Get(developerAppGrantGetReq models.DeveloperAppGrantGetRequestModel) (models.DeveloperAppGrantGetResponseModel, error)
}

type DeveloperApplicationGrantService struct {
	GormClient *gorm.DB
}

func (d *DeveloperApplicationGrantService) Create(developerAppGrantCreateReq models.DeveloperAppGrantCreateRequestModel) (models.DeveloperAppGrantCreateResponseModel, error) {
	newDeveloperAppGrant := entities.DeveloperApplicationGrant{
		BaseDetails: entities.BaseEntity{
			Id: uuid.New().String(),
		},
		PlayerUsername: developerAppGrantCreateReq.PlayerUsername,
		ApplicationId:  developerAppGrantCreateReq.ApplicationId,
		ExpiresAt:      developerAppGrantCreateReq.ExpiresAt,
		Scope:          developerAppGrantCreateReq.Scope,
	}
	response := d.GormClient.Create(&newDeveloperAppGrant)
	if response.Error != nil {
		log.Println("error creating dev-app-grant")
		log.Println(response.Error.Error())
		return *new(models.DeveloperAppGrantCreateResponseModel), nil
	}

	log.Println("create dev-app-grant")
	return models.DeveloperAppGrantCreateResponseModel{
		DeveloperAppGrantId: newDeveloperAppGrant.BaseDetails.Id,
		Username:            newDeveloperAppGrant.PlayerUsername,
		ApplicationId:       newDeveloperAppGrant.ApplicationId,
		Scope:               newDeveloperAppGrant.Scope,
		ExpiresAt:           newDeveloperAppGrant.ExpiresAt,
	}, nil
}

func (d *DeveloperApplicationGrantService) Get(developerAppGrantGetReq models.DeveloperAppGrantGetRequestModel) (models.DeveloperAppGrantGetResponseModel, error) {
	targetGrant := entities.DeveloperApplicationGrant{
		BaseDetails: entities.BaseEntity{
			Id: developerAppGrantGetReq.DeveloperAppGrantId,
		},
	}
	gormResponse := d.GormClient.First(&targetGrant)
	if gormResponse.Error != nil {
		log.Println("error not found")
		return *new(models.DeveloperAppGrantGetResponseModel), nil
	}

	log.Println("successfully found grant")
	return models.DeveloperAppGrantGetResponseModel{
		DeveloperAppGrantIdId: targetGrant.BaseDetails.Id,
		PlayerUsername:        targetGrant.PlayerUsername,
		ApplicationId:         targetGrant.ApplicationId,
		Scope:                 targetGrant.Scope,
		ExpiresAt:             targetGrant.ExpiresAt,
	}, nil
}

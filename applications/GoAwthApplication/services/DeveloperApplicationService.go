package services

import (
	"github.com/google/uuid"
	"github.com/mrkresnofatih/go-awth/entities"
	"github.com/mrkresnofatih/go-awth/models"
	"gorm.io/gorm"
	"log"
)

type IDeveloperApplicationService interface {
	Create(createRequest models.DeveloperApplicationCreateRequestModel) (models.DeveloperApplicationGetResponseModel, error)
	List(listRequest models.DeveloperApplicationListRequestModel) (models.DeveloperApplicationListResponseModel, error)
	Get(getRequest models.DeveloperApplicationGetRequestModel) (models.DeveloperApplicationGetResponseModel, error)
}

type DeveloperApplicationService struct {
	GormClient *gorm.DB
}

func (d *DeveloperApplicationService) Create(createRequest models.DeveloperApplicationCreateRequestModel) (models.DeveloperApplicationGetResponseModel, error) {
	newDevApp := entities.DeveloperApplication{
		BaseDetails: entities.BaseEntity{
			Id: uuid.New().String(),
		},
		DeveloperName:      createRequest.DeveloperName,
		Name:               createRequest.Name,
		LogoUrl:            createRequest.LogoUrl,
		Secret:             uuid.New().String(),
		SuccessRedirectUri: createRequest.SuccessRedirectUri,
		FailedRedirectUri:  createRequest.FailedRedirectUri,
	}
	response := d.GormClient.Create(&newDevApp)
	if response.Error != nil {
		log.Println("failed to save dev app to db")
		log.Println(response.Error.Error())
		return *new(models.DeveloperApplicationGetResponseModel), nil
	}

	log.Println("dev app created!")
	return models.DeveloperApplicationGetResponseModel{
		DeveloperApplicationId: newDevApp.BaseDetails.Id,
		DeveloperName:          newDevApp.DeveloperName,
		LogoUrl:                newDevApp.LogoUrl,
		Name:                   newDevApp.Name,
		Secret:                 newDevApp.Secret,
		SuccessRedirectUri:     newDevApp.SuccessRedirectUri,
		FailedRedirectUri:      newDevApp.FailedRedirectUri,
	}, nil
}

func (d *DeveloperApplicationService) List(listRequest models.DeveloperApplicationListRequestModel) (models.DeveloperApplicationListResponseModel, error) {
	targetDeveloper := entities.Developer{
		DeveloperName: listRequest.DeveloperName,
	}
	gormResponse := d.GormClient.Model(&entities.Developer{}).Preload("DeveloperApplications", func(db *gorm.DB) *gorm.DB {
		return db.Offset(listRequest.PageSize * (listRequest.Page - 1)).Limit(listRequest.PageSize)
	}).First(&targetDeveloper)
	if gormResponse.Error != nil {
		log.Println("find developer w/ paginated dev apps failed")
		return *new(models.DeveloperApplicationListResponseModel), nil
	}

	log.Println(len(targetDeveloper.DeveloperApplications))

	var apps []models.DeveloperApplicationGetResponseModel
	for _, app := range targetDeveloper.DeveloperApplications {
		apps = append(apps, models.DeveloperApplicationGetResponseModel{
			DeveloperApplicationId: app.BaseDetails.Id,
			DeveloperName:          app.DeveloperName,
			LogoUrl:                app.LogoUrl,
			Name:                   app.Name,
			Secret:                 app.Secret,
			SuccessRedirectUri:     app.SuccessRedirectUri,
			FailedRedirectUri:      app.FailedRedirectUri,
		})
	}

	log.Println("find developer w/ paginated dev apps successful")
	return models.DeveloperApplicationListResponseModel{
		DeveloperName: listRequest.DeveloperName,
		Applications:  apps,
	}, nil
}

func (d *DeveloperApplicationService) Get(getRequest models.DeveloperApplicationGetRequestModel) (models.DeveloperApplicationGetResponseModel, error) {
	//TODO implement me
	panic("implement me")
}

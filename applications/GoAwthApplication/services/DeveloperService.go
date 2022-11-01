package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrkresnofatih/go-awth/entities"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"gorm.io/gorm"
	"log"
	"time"
)

type IDeveloperService interface {
	Signup(developerSignupRequest models.DeveloperSignupRequestModel) (models.DeveloperSignupResponseModel, error)
	Login(developerLoginRequest models.DeveloperLoginRequestModel) (models.DeveloperLoginResponseModel, error)
}

type DeveloperService struct {
	GormClient *gorm.DB
}

func (d *DeveloperService) Signup(developerSignupRequest models.DeveloperSignupRequestModel) (models.DeveloperSignupResponseModel, error) {
	newDeveloper := entities.Developer{
		BaseDetails: entities.BaseEntity{
			Id: uuid.New().String(),
		},
		DeveloperName: developerSignupRequest.DeveloperName,
		Password:      developerSignupRequest.Password,
	}
	response := d.GormClient.Create(&newDeveloper)
	if response.Error != nil {
		log.Println("error creating developer")
		log.Println(response.Error.Error())
		return *new(models.DeveloperSignupResponseModel), errors.New("failed to create developer")
	}

	log.Println("developer signup success w/ new id: " + newDeveloper.BaseDetails.Id)
	return models.DeveloperSignupResponseModel{
		DeveloperName: newDeveloper.DeveloperName,
	}, nil
}

func (d *DeveloperService) Login(developerLoginRequest models.DeveloperLoginRequestModel) (models.DeveloperLoginResponseModel, error) {
	targetDeveloper := entities.Developer{
		DeveloperName: developerLoginRequest.DeveloperName,
	}
	gormResponse := d.
		GormClient.
		Where(&entities.Developer{DeveloperName: developerLoginRequest.DeveloperName}).
		First(&targetDeveloper)
	if gormResponse.Error != nil {
		log.Println(gormResponse.Error)
		return *new(models.DeveloperLoginResponseModel), errors.New("failed to find developer")
	}

	if targetDeveloper.Password != developerLoginRequest.Password {
		log.Println("developer password incorrect")
		return *new(models.DeveloperLoginResponseModel), errors.New("incorrect developer password")
	}

	basicJwtBuilder := &jwt.BasicJwtTokenBuilder{
		ExpiresAfter: time.Hour * 1,
	}
	developerNameJwtBuilder := &jwt.DeveloperNameJwtTokenBuilder{
		JwtTokenBuilder: basicJwtBuilder,
		DeveloperName:   targetDeveloper.DeveloperName,
	}
	token, err := developerNameJwtBuilder.Build()
	if err != nil {
		log.Println("error creating dev token")
		return *new(models.DeveloperLoginResponseModel), errors.New("failed creating developer token")
	}

	log.Println("dev login success")
	return models.DeveloperLoginResponseModel{
		DeveloperName: targetDeveloper.DeveloperName,
		AccessToken:   token,
	}, nil
}

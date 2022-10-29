package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrkresnofatih/go-awth/entities"
	"github.com/mrkresnofatih/go-awth/models"
	"gorm.io/gorm"
	"log"
)

type IPlayerService interface {
	Signup(playerSignupRequest models.PlayerSignupRequestModel) (models.PlayerSignupResponseModel, error)
	Login(playerLoginRequest models.PlayerLoginRequestModel) (models.PlayerLoginResponseModel, error)
	Update(playerUpdateRequest models.PlayerUpdateRequestModel) (models.PlayerUpdateResponseModel, error)
	Get(playerGetRequest models.PlayerGetRequestModel) (models.PlayerGetResponseModel, error)
}

type PlayerService struct {
	GormClient *gorm.DB
}

func (p PlayerService) Get(playerGetRequest models.PlayerGetRequestModel) (models.PlayerGetResponseModel, error) {
	targetPlayer := entities.PlayerEntity{
		Username: playerGetRequest.Username,
	}
	err := p.GormClient.First(&targetPlayer)
	if err != nil {
		log.Println("error not found")
		return *new(models.PlayerGetResponseModel), errors.New("player not found")
	}

	log.Println("Successfully found player")
	return models.PlayerGetResponseModel{
		Username: targetPlayer.Username,
		FullName: targetPlayer.FullName,
		ImageUrl: targetPlayer.ImageUrl,
	}, nil
}

func (p PlayerService) Signup(playerSignupRequest models.PlayerSignupRequestModel) (models.PlayerSignupResponseModel, error) {
	newUser := entities.PlayerEntity{
		Username: playerSignupRequest.Username,
		FullName: playerSignupRequest.FullName,
		ImageUrl: playerSignupRequest.ImageUrl,
		Password: playerSignupRequest.Password,
		BaseDetails: entities.BaseEntity{
			Id: uuid.New().String(),
		},
	}
	response := p.GormClient.Create(&newUser)
	if response.Error != nil {
		log.Println("error creating user")
		return *new(models.PlayerSignupResponseModel), nil
	}

	log.Println("player signup success w/ new id: " + newUser.BaseDetails.Id)
	return models.PlayerSignupResponseModel{
		Username: newUser.Username,
		FullName: newUser.FullName,
	}, nil
}

func (p PlayerService) Login(playerLoginRequest models.PlayerLoginRequestModel) (models.PlayerLoginResponseModel, error) {
	targetPlayer := entities.PlayerEntity{
		Username: playerLoginRequest.Username,
	}
	err := p.GormClient.First(&targetPlayer)
	if err != nil {
		log.Println("error target player not found")
		return *new(models.PlayerLoginResponseModel), errors.New("failed To Find Player")
	}

	if targetPlayer.Password != playerLoginRequest.Password {
		log.Println("password incorrect")
		return *new(models.PlayerLoginResponseModel), errors.New("incorrect password")
	}

	log.Println("login success")
	return models.PlayerLoginResponseModel{
		Username:    targetPlayer.Username,
		AccessToken: "access-token",
	}, nil
}

func (p PlayerService) Update(playerUpdateRequest models.PlayerUpdateRequestModel) (models.PlayerUpdateResponseModel, error) {
	targetPlayer := entities.PlayerEntity{
		Username: playerUpdateRequest.Username,
	}
	err := p.GormClient.First(&targetPlayer)
	if err != nil {
		log.Println("player for update not found")
		return *new(models.PlayerUpdateResponseModel), nil
	}

	targetPlayer.FullName = playerUpdateRequest.FullName
	targetPlayer.Password = playerUpdateRequest.Password
	targetPlayer.ImageUrl = playerUpdateRequest.ImageUrl

	response := p.GormClient.Save(&targetPlayer)
	if response.Error != nil {
		log.Println("error saving target player")
		return *new(models.PlayerUpdateResponseModel), nil
	}

	log.Println("finish updating target player")
	return models.PlayerUpdateResponseModel{
		Username: targetPlayer.Username,
		FullName: targetPlayer.FullName,
		ImageUrl: targetPlayer.ImageUrl,
	}, nil
}

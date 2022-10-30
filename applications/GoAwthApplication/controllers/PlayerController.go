package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/endpoints/playerEndpoints"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
)

type PlayerController struct {
	PlayerService services.IPlayerService
}

func (p *PlayerController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/player",
	}

	playerSignupEndpoint := &playerEndpoints.PlayerSignupEndpoint{
		PlayerService: p.PlayerService,
	}
	playerSignupWithValidation := &RequireValidationDecorator[models.PlayerSignupRequestModel]{
		Endpoint: playerSignupEndpoint,
	}
	controllerRouter.AddEndpoint(playerSignupWithValidation)

	playerSignupValidationEndpoint := &playerEndpoints.PlayerSignupValidationEndpoint{}
	controllerRouter.AddEndpoint(playerSignupValidationEndpoint)

	playerLoginEndpoint := &playerEndpoints.PlayerLoginEndpoint{
		PlayerService: p.PlayerService,
	}
	playerLoginWithValidation := &RequireValidationDecorator[models.PlayerLoginRequestModel]{
		Endpoint: playerLoginEndpoint,
	}
	controllerRouter.AddEndpoint(playerLoginWithValidation)

	playerGetProfileEndpoint := &playerEndpoints.PlayerGetMyProfileEndpoint{
		PlayerService: p.PlayerService,
	}
	playerGetProfileWithAuth := &RequireAuthenticationDecorator{
		Endpoint: playerGetProfileEndpoint,
	}
	controllerRouter.AddEndpoint(playerGetProfileWithAuth)

	controllerRouter.Build()
}

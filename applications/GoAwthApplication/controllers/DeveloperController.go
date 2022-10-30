package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/endpoints/developerEndpoints"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
)

type DeveloperController struct {
	DeveloperService services.IDeveloperService
}

func (d *DeveloperController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/developer",
	}

	devSignupEndpoint := &developerEndpoints.DeveloperSignupEndpoint{
		DeveloperService: d.DeveloperService,
	}
	devSignupWithValidation := &RequireValidationDecorator[models.DeveloperSignupRequestModel]{
		Endpoint: devSignupEndpoint,
	}
	controllerRouter.AddEndpoint(devSignupWithValidation)

	devLoginEndpoint := &developerEndpoints.DeveloperLoginEndpoint{
		DeveloperService: d.DeveloperService,
	}
	devLoginWithValidation := &RequireValidationDecorator[models.DeveloperLoginRequestModel]{
		Endpoint: devLoginEndpoint,
	}
	controllerRouter.AddEndpoint(devLoginWithValidation)

	controllerRouter.Build()
}

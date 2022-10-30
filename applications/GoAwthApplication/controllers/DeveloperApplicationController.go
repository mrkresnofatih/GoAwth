package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/endpoints/devappsEndpoints"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
)

type DeveloperApplicationController struct {
	DeveloperApplicationService services.IDeveloperApplicationService
}

func (d *DeveloperApplicationController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/dev-apps",
	}

	createDevApplicationEndpoint := &devappsEndpoints.DeveloperApplicationCreateEndpoint{
		DeveloperApplicationService: d.DeveloperApplicationService,
	}
	createDevAppWithValidation := &RequireValidationDecorator[models.DeveloperApplicationCreateRequestModel]{
		Endpoint: createDevApplicationEndpoint,
	}
	createDevAppWithDevAuth := &RequireDeveloperAccessDecorator{
		Endpoint: createDevAppWithValidation,
	}
	controllerRouter.AddEndpoint(createDevAppWithDevAuth)

	listDevAppsEndpoint := &devappsEndpoints.DeveloperApplicationListEndpoint{
		DeveloperApplicationService: d.DeveloperApplicationService,
	}
	listDevAppWithValidation := &RequireValidationDecorator[models.DeveloperApplicationListRequestModel]{
		Endpoint: listDevAppsEndpoint,
	}
	listDevAppWithDevAuth := &RequireDeveloperAccessDecorator{
		Endpoint: listDevAppWithValidation,
	}
	controllerRouter.AddEndpoint(listDevAppWithDevAuth)

	controllerRouter.Build()
}

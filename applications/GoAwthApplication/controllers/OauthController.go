package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/endpoints/oauthEndpoints"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
)

type OauthController struct {
	OauthService services.IOauthService
}

func (o *OauthController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/oauth2",
	}

	getConsentEndpoint := &oauthEndpoints.OauthGetConsentEndpoint{
		OauthService: o.OauthService,
	}
	getConsentWithValidation := &RequireValidationDecorator[models.OauthGetConsentRequestModel]{
		Endpoint: getConsentEndpoint,
	}
	getConsentWithPlayerRole := &RequirePlayerAccessDecorator{
		Endpoint: getConsentWithValidation,
	}
	controllerRouter.AddEndpoint(getConsentWithPlayerRole)

	agreeConsentEndpoint := &oauthEndpoints.OauthAgreeConsentEndpoint{
		OauthService: o.OauthService,
	}
	agreeConsentWithValidation := &RequireValidationDecorator[models.OauthAgreeConsentRequestModel]{
		Endpoint: agreeConsentEndpoint,
	}
	agreeConsentWithPlayerRole := &RequirePlayerAccessDecorator{
		Endpoint: agreeConsentWithValidation,
	}
	controllerRouter.AddEndpoint(agreeConsentWithPlayerRole)

	rejectConsentEndpoint := &oauthEndpoints.OauthRejectConsentEndpoint{
		OauthService: o.OauthService,
	}
	rejectConsentWithValidation := &RequireValidationDecorator[models.OauthRejectConsentResponseModel]{
		Endpoint: rejectConsentEndpoint,
	}
	rejectConsentWithPlayerRole := &RequirePlayerAccessDecorator{
		Endpoint: rejectConsentWithValidation,
	}
	controllerRouter.AddEndpoint(rejectConsentWithPlayerRole)

	authGrantEndpoint := &oauthEndpoints.OauthAuthenticateGrantEndpoint{
		OauthService: o.OauthService,
	}
	authGrantWithValidation := &RequireValidationDecorator[models.OauthAuthenticateGrantRequestModel]{
		Endpoint: authGrantEndpoint,
	}
	controllerRouter.AddEndpoint(authGrantWithValidation)

	controllerRouter.Build()
}

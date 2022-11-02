package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth-client-app/endpoints"
	"github.com/mrkresnofatih/go-awth-client-app/services"
)

type AuthController struct {
	AuthService services.IAuthService
}

func (a AuthController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/auth",
	}

	authAuthenticateEndpoint := &endpoints.AuthAuthenticateEndpoint{
		AuthService: a.AuthService,
	}
	controllerRouter.AddEndpoint(authAuthenticateEndpoint)

	authLoginEndpoint := &endpoints.AuthLoginEndpoint{}
	controllerRouter.AddEndpoint(authLoginEndpoint)

	controllerRouter.Build()
}

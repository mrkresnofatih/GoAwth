package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IServer interface {
	Initialize()
	AddController(controller IController)
}

type ApplicationServer struct {
	Router      *echo.Echo
	Controllers []IController
}

func (a *ApplicationServer) Initialize() {
	a.Router = echo.New()
	a.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
	}))
	for _, controller := range a.Controllers {
		controller.Register(a.Router)
	}
}

func (a *ApplicationServer) AddController(controller IController) {
	a.Controllers = append(a.Controllers, controller)
}

type IController interface {
	Register(echo *echo.Echo)
}

type IEndpoint interface {
	GetHandler() echo.HandlerFunc
	GetMethod() string
	GetPath() string
	Register(group *echo.Group)
}

type IRouter interface {
	Build()
	AddEndpoint(endpoint IEndpoint)
}

type ControllerRouter struct {
	MainRouter *echo.Echo
	Router     *echo.Group
	PathPrefix string
	Endpoints  []IEndpoint
}

func (c *ControllerRouter) Build() {
	c.Router = c.MainRouter.Group(c.PathPrefix)
	for _, endpoint := range c.Endpoints {
		endpoint.Register(c.Router)
	}
}

func (c *ControllerRouter) AddEndpoint(endpoint IEndpoint) {
	c.Endpoints = append(c.Endpoints, endpoint)
}

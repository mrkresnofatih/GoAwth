package controllers

import (
	"bytes"
	"encoding/json"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"io"
	"log"
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

type RequireAuthenticationDecorator struct {
	Endpoint IEndpoint
}

func (r *RequireAuthenticationDecorator) GetPath() string {
	return r.Endpoint.GetPath()
}

func (r *RequireAuthenticationDecorator) GetMethod() string {
	return r.Endpoint.GetMethod()
}

func (r *RequireAuthenticationDecorator) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) < 7 {
			log.Println("auth header not valid")
			return models.SendBadResponse(c, "auth failed")
		}

		jwtToken := authHeader[7:] // removes bearer
		isJwtTokenValid := jwt.GetValidityFromToken(jwtToken)
		if !isJwtTokenValid {
			return models.SendBadResponse(c, "auth failed")
		}
		return r.Endpoint.GetHandler()(c)
	}
}

func (r *RequireAuthenticationDecorator) Register(group *echo.Group) {
	group.Match([]string{r.GetMethod()}, r.GetPath(), r.GetHandler())
}

type RequireDeveloperAccessDecorator struct {
	Endpoint IEndpoint
}

func (r *RequireDeveloperAccessDecorator) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) < 7 {
			log.Println("auth header not valid")
			return models.SendBadResponse(c, "auth failed")
		}

		jwtToken := authHeader[7:] // removes bearer
		isJwtTokenValid := jwt.GetValidityFromToken(jwtToken)
		if !isJwtTokenValid {
			return models.SendBadResponse(c, "auth failed")
		}

		developerName, err := jwt.GetClaimFromToken[string](jwtToken, jwt.ApplicationJwtClaimsKeyDeveloperName)
		if err != nil {
			log.Println("developerName claim not found")
			return models.SendBadResponse(c, "Invalid developer access token")
		}

		log.Println("developer access granted: " + developerName)
		return r.Endpoint.GetHandler()(c)
	}
}

func (r *RequireDeveloperAccessDecorator) GetMethod() string {
	return r.Endpoint.GetMethod()
}

func (r *RequireDeveloperAccessDecorator) GetPath() string {
	return r.Endpoint.GetPath()
}

func (r *RequireDeveloperAccessDecorator) Register(group *echo.Group) {
	group.Match([]string{r.GetMethod()}, r.GetPath(), r.GetHandler())
}

type RequireValidationDecorator[T interface{}] struct {
	Endpoint IEndpoint
}

func (r *RequireValidationDecorator[T]) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			log.Println("cannot read request body")
			return models.SendBadResponse(c, "Failed to read req body")
		}

		var bodyData T
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			log.Println("json parse failed")
			return models.SendBadResponse(c, "Failed to json parse")
		}

		validator := validator2.New()
		err = validator.Struct(bodyData)
		if err != nil {
			if _, ok := err.(*validator2.InvalidValidationError); ok {
				log.Println(err)
				return models.SendBadResponse(c, "Invalid validation error")
			}

			errors := err.(validator2.ValidationErrors)
			log.Println(errors)

			return models.SendBadResponse(c, "Req Validation Errors")
		}

		newR := c.Request().Clone(c.Request().Context())
		c.Request().Body = io.NopCloser(bytes.NewReader(body))
		newR.Body = io.NopCloser(bytes.NewReader(body))
		err = c.Request().ParseForm()
		if err != nil {
			log.Println("Error cloning request")
			return models.SendBadResponse(c, "Failed to duplicate request")
		}
		c.SetRequest(newR)
		return r.Endpoint.GetHandler()(c)
	}
}

func (r *RequireValidationDecorator[T]) GetMethod() string {
	return r.Endpoint.GetMethod()
}

func (r *RequireValidationDecorator[T]) GetPath() string {
	return r.Endpoint.GetPath()
}

func (r *RequireValidationDecorator[T]) Register(group *echo.Group) {
	group.Match([]string{r.GetMethod()}, r.GetPath(), r.GetHandler())
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

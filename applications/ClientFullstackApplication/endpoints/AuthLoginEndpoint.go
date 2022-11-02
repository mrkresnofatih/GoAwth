package endpoints

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthLoginEndpoint struct{}

func (a *AuthLoginEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<a href=\"http://localhost:3000/oauth2?applicationid=043af8b1-914c-45e5-a01e-cef860ed6875&granttype=grantid&scopes=openid\">login</a>")
	}
}

func (a *AuthLoginEndpoint) GetMethod() string {
	return http.MethodGet
}

func (a *AuthLoginEndpoint) GetPath() string {
	return "/login"
}

func (a *AuthLoginEndpoint) Register(group *echo.Group) {
	group.Match([]string{a.GetMethod()}, a.GetPath(), a.GetHandler())
}

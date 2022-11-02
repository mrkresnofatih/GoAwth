package endpoints

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth-client-app/models"
	"github.com/mrkresnofatih/go-awth-client-app/services"
)

type AuthAuthenticateEndpoint struct {
	AuthService services.IAuthService
}

func (a *AuthAuthenticateEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		grantId := c.QueryParams().Get("grantid")
		if len(grantId) == 0 {
			return c.HTML(http.StatusBadRequest, "<p>No grant id provided</p>")
		}

		authGrantResponse, err := a.AuthService.AuthenticateGrant(models.AuthAuthenticateGrantRequestModel{
			GrantId:           grantId,
			ApplicationId:     "043af8b1-914c-45e5-a01e-cef860ed6875", // GoAwth AppId
			ApplicationSecret: "cec274c5-a6e5-4055-940f-bd3173833438", // GoAwth secret
		})
		if err != nil {
			log.Println("auth grant failed")
			return c.HTML(http.StatusBadRequest, "<p>Auth failed</p>")
		}

		openIdResponse, err := a.AuthService.GetOpenId(models.AuthOpenIdRequestModel{
			Token: authGrantResponse.GrantToken,
		})
		if err != nil {
			log.Println("openid connect failed")
			return c.HTML(http.StatusBadRequest, "<p>OpenID Connect Failed</p>")
		}

		return c.JSON(http.StatusOK, openIdResponse)
	}
}

func (a *AuthAuthenticateEndpoint) GetMethod() string {
	return http.MethodGet
}

func (a *AuthAuthenticateEndpoint) GetPath() string {
	return "/authenticate"
}

func (a *AuthAuthenticateEndpoint) Register(group *echo.Group) {
	group.Match([]string{a.GetMethod()}, a.GetPath(), a.GetHandler())
}

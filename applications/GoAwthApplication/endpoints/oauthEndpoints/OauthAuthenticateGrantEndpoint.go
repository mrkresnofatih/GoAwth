package oauthEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type OauthAuthenticateGrantEndpoint struct {
	OauthService services.IOauthService
}

func (o *OauthAuthenticateGrantEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		authenticateGrantReq := new(models.OauthAuthenticateGrantRequestModel)
		_ = c.Bind(authenticateGrantReq)

		authGrantResp, err := o.OauthService.AuthenticateGrant(*authenticateGrantReq)
		if err != nil {
			log.Println("error failed to authenticate grant")
			return models.SendBadResponse(c, "failed to authenticate grant")
		}

		log.Println("authenticate grant success")
		return models.SendGoodResponse[models.OauthAuthenticateGrantResponseModel](c, authGrantResp)
	}
}

func (o *OauthAuthenticateGrantEndpoint) GetMethod() string {
	return http.MethodPost
}

func (o *OauthAuthenticateGrantEndpoint) GetPath() string {
	return "/authenticate-grant"
}

func (o *OauthAuthenticateGrantEndpoint) Register(group *echo.Group) {
	group.Match([]string{o.GetMethod()}, o.GetPath(), o.GetHandler())
}

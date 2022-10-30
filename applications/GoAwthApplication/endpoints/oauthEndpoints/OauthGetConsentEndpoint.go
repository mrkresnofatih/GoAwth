package oauthEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"net/http"
)

type OauthGetConsentEndpoint struct {
	OauthService services.IOauthService
}

func (o OauthGetConsentEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		getConsentReq := new(models.OauthGetConsentRequestModel)
		_ = c.Bind(getConsentReq)

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		username, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyUsername)

		getConsentReq.PlayerUsername = username

		getConsentData, err := o.OauthService.GetConsent(*getConsentReq)
		if err != nil {
			log.Println("failed to get consent")
			return models.SendBadResponse(c, "Failed to get consent")
		}
		
		log.Println("get consent success")
		return models.SendGoodResponse[models.OauthGetConsentResponseModel](c, getConsentData)
	}
}

func (o OauthGetConsentEndpoint) GetMethod() string {
	return http.MethodPost
}

func (o OauthGetConsentEndpoint) GetPath() string {
	return "/get-consent"
}

func (o OauthGetConsentEndpoint) Register(group *echo.Group) {
	group.Match([]string{o.GetMethod()}, o.GetPath(), o.GetHandler())
}

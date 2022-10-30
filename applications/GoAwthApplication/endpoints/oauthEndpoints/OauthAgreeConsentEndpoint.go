package oauthEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"net/http"
)

type OauthAgreeConsentEndpoint struct {
	OauthService services.IOauthService
}

func (o *OauthAgreeConsentEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		agreeConsentReq := new(models.OauthAgreeConsentRequestModel)
		_ = c.Bind(agreeConsentReq)

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		username, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyUsername)

		agreeConsentReq.PlayerUsername = username

		agreeConsentResp, err := o.OauthService.AgreeConsent(*agreeConsentReq)
		if err != nil {
			log.Println("error failed to agree consent")
			return models.SendBadResponse(c, "Agree to consent failed")
		}

		log.Println("agree consent success")
		return models.SendGoodResponse[models.OauthAgreeConsentResponseModel](c, agreeConsentResp)
	}
}

func (o *OauthAgreeConsentEndpoint) GetMethod() string {
	return http.MethodPost
}

func (o *OauthAgreeConsentEndpoint) GetPath() string {
	return "/agree-consent"
}

func (o *OauthAgreeConsentEndpoint) Register(group *echo.Group) {
	group.Match([]string{o.GetMethod()}, o.GetPath(), o.GetHandler())
}

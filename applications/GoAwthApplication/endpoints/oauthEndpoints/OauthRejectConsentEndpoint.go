package oauthEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type OauthRejectConsentEndpoint struct {
	OauthService services.IOauthService
}

func (o *OauthRejectConsentEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		rejectConsentReq := new(models.OauthRejectConsentRequestModel)
		_ = c.Bind(rejectConsentReq)

		rejectConsentResponse, err := o.OauthService.RejectConsent(*rejectConsentReq)
		if err != nil {
			log.Println("failed to get consent")
			return models.SendBadResponse(c, "Failed to get consent")
		}

		log.Println("get consent success")
		return models.SendGoodResponse[models.OauthRejectConsentResponseModel](c, rejectConsentResponse)
	}
}

func (o *OauthRejectConsentEndpoint) GetMethod() string {
	return http.MethodPost
}

func (o *OauthRejectConsentEndpoint) GetPath() string {
	return "/reject-consent"
}

func (o *OauthRejectConsentEndpoint) Register(group *echo.Group) {
	group.Match([]string{o.GetMethod()}, o.GetPath(), o.GetHandler())
}

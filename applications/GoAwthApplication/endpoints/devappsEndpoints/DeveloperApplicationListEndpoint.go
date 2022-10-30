package devappsEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"net/http"
)

type DeveloperApplicationListEndpoint struct {
	DeveloperApplicationService services.IDeveloperApplicationService
}

func (d *DeveloperApplicationListEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		listReq := new(models.DeveloperApplicationListRequestModel)
		_ = c.Bind(listReq)

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		developerName, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyDeveloperName)

		listReq.DeveloperName = developerName
		listDevApp, err := d.DeveloperApplicationService.List(*listReq)
		if err != nil {
			log.Println("failed to list dev apps")
			return models.SendBadResponse(c, "Failed to list dev apps")
		}

		return models.SendGoodResponse[models.DeveloperApplicationListResponseModel](c, listDevApp)
	}
}

func (d *DeveloperApplicationListEndpoint) GetMethod() string {
	return http.MethodPost
}

func (d *DeveloperApplicationListEndpoint) GetPath() string {
	return "/list"
}

func (d *DeveloperApplicationListEndpoint) Register(group *echo.Group) {
	group.Match([]string{d.GetMethod()}, d.GetPath(), d.GetHandler())
}

package devappsEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"net/http"
)

type DeveloperApplicationCreateEndpoint struct {
	DeveloperApplicationService services.IDeveloperApplicationService
}

func (d *DeveloperApplicationCreateEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		createReq := new(models.DeveloperApplicationCreateRequestModel)
		_ = c.Bind(createReq)

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		developerName, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyDeveloperName)

		createReq.DeveloperName = developerName
		createdDevApp, err := d.DeveloperApplicationService.Create(*createReq)
		if err != nil {
			log.Println("failed to create dev application")
			return models.SendBadResponse(c, "Failed to create developer application")
		}

		return models.SendGoodResponse[models.DeveloperApplicationGetResponseModel](c, createdDevApp)
	}
}

func (d *DeveloperApplicationCreateEndpoint) GetMethod() string {
	return http.MethodPost
}

func (d *DeveloperApplicationCreateEndpoint) GetPath() string {
	return "/create"
}

func (d *DeveloperApplicationCreateEndpoint) Register(group *echo.Group) {
	group.Match([]string{d.GetMethod()}, d.GetPath(), d.GetHandler())
}

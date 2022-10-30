package developerEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type DeveloperLoginEndpoint struct {
	DeveloperService services.IDeveloperService
}

func (d *DeveloperLoginEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginReq := new(models.DeveloperLoginRequestModel)
		_ = c.Bind(loginReq)

		devLoginResponse, err := d.DeveloperService.Login(*loginReq)
		if err != nil {
			log.Println("failed to login developer account")
			return models.SendBadResponse(c, "Failed to Login")
		}

		return models.SendGoodResponse[models.DeveloperLoginResponseModel](c, devLoginResponse)
	}
}

func (d *DeveloperLoginEndpoint) GetMethod() string {
	return http.MethodPost
}

func (d *DeveloperLoginEndpoint) GetPath() string {
	return "/login"
}

func (d *DeveloperLoginEndpoint) Register(group *echo.Group) {
	group.Match([]string{d.GetMethod()}, d.GetPath(), d.GetHandler())
}

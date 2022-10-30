package developerEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type DeveloperSignupEndpoint struct {
	DeveloperService services.IDeveloperService
}

func (d *DeveloperSignupEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		signupReq := new(models.DeveloperSignupRequestModel)
		_ = c.Bind(signupReq)

		createdDeveloperAccount, err := d.DeveloperService.Signup(*signupReq)
		if err != nil {
			log.Println("failed to create developer account")
			return models.SendBadResponse(c, "Failed to signup")
		}

		return models.SendGoodResponse[models.DeveloperSignupResponseModel](c, createdDeveloperAccount)
	}
}

func (d *DeveloperSignupEndpoint) GetMethod() string {
	return http.MethodPost
}

func (d *DeveloperSignupEndpoint) GetPath() string {
	return "/sign-up"
}

func (d *DeveloperSignupEndpoint) Register(group *echo.Group) {
	group.Match([]string{d.GetMethod()}, d.GetPath(), d.GetHandler())
}

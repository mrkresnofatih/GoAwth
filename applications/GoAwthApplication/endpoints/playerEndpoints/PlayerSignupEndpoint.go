package playerEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type PlayerSignupEndpoint struct {
	PlayerService services.IPlayerService
}

func (p *PlayerSignupEndpoint) GetMethod() string {
	return http.MethodPost
}

func (p *PlayerSignupEndpoint) GetPath() string {
	return "/sign-up"
}

func (p *PlayerSignupEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		signupRequest := new(models.PlayerSignupRequestModel)
		if err := c.Bind(signupRequest); err != nil {
			log.Println("Failed to bind request body")
			return models.SendBadResponse(c, "Failed to bind request body")
		}

		createdPlayer, err := p.PlayerService.Signup(*signupRequest)
		if err != nil {
			log.Println("Failed to signup")
			return models.SendBadResponse(c, "Failed to signup")
		}

		return models.SendGoodResponse[models.PlayerSignupResponseModel](c, createdPlayer)
	}
}

func (p *PlayerSignupEndpoint) Register(group *echo.Group) {
	group.Match([]string{p.GetMethod()}, p.GetPath(), p.GetHandler())
}

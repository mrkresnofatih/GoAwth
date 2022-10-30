package playerEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"log"
	"net/http"
)

type PlayerLoginEndpoint struct {
	PlayerService services.IPlayerService
}

func (p *PlayerLoginEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginReq := new(models.PlayerLoginRequestModel)
		_ = c.Bind(loginReq)

		loginResponse, err := p.PlayerService.Login(*loginReq)
		if err != nil {
			log.Println("Failed to login")
			return models.SendBadResponse(c, err.Error())
		}
		return models.SendGoodResponse[models.PlayerLoginResponseModel](c, loginResponse)
	}
}

func (p *PlayerLoginEndpoint) GetMethod() string {
	return http.MethodPost
}

func (p *PlayerLoginEndpoint) GetPath() string {
	return "/login"
}

func (p *PlayerLoginEndpoint) Register(group *echo.Group) {
	group.Match([]string{p.GetMethod()}, p.GetPath(), p.GetHandler())
}

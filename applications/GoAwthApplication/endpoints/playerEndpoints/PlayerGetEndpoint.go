package playerEndpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/services"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"net/http"
)

type PlayerGetMyProfileEndpoint struct {
	PlayerService services.IPlayerService
}

func (p *PlayerGetMyProfileEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		username, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyUsername)

		getPlayerResponse, err := p.PlayerService.Get(models.PlayerGetRequestModel{
			Username: username,
		})
		if err != nil {
			log.Println("get player failed")
			return models.SendBadResponse(c, "failed to get profile")
		}
		return models.SendGoodResponse[models.PlayerGetResponseModel](c, getPlayerResponse)
	}
}

func (p *PlayerGetMyProfileEndpoint) GetMethod() string {
	return http.MethodGet
}

func (p *PlayerGetMyProfileEndpoint) GetPath() string {
	return "/get-my-profile"
}

func (p *PlayerGetMyProfileEndpoint) Register(group *echo.Group) {
	group.Match([]string{p.GetMethod()}, p.GetPath(), p.GetHandler())
}

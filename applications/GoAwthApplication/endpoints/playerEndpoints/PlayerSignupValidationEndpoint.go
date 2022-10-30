package playerEndpoints

import (
	validator2 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/go-awth/models"
	"log"
	"net/http"
)

type PlayerSignupValidationEndpoint struct {
}

func (p *PlayerSignupValidationEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		signupRequest := new(models.PlayerSignupRequestModel)
		if err := c.Bind(signupRequest); err != nil {
			log.Println("error binding req body")
			return models.SendBadResponse(c, "error binding req body")
		}

		validator := validator2.New()
		err := validator.Struct(signupRequest)
		if err != nil {
			if _, ok := err.(*validator2.InvalidValidationError); ok {
				log.Println(err)
				return models.SendBadResponse(c, "Invalid validation error")
			}

			errors := err.(validator2.ValidationErrors)
			var errorMessages []string
			for _, validationError := range errors {
				errorMessages = append(errorMessages, validationError.Field())
			}

			return models.SendGoodResponse[[]string](c, errorMessages)
		}

		return models.SendGoodResponse[[]string](c, []string{})
	}
}

func (p *PlayerSignupValidationEndpoint) GetMethod() string {
	return http.MethodPost
}

func (p *PlayerSignupValidationEndpoint) GetPath() string {
	return "/validate-sign-up"
}

func (p *PlayerSignupValidationEndpoint) Register(group *echo.Group) {
	group.Match([]string{p.GetMethod()}, p.GetPath(), p.GetHandler())
}

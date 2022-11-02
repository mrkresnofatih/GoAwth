package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrkresnofatih/go-awth-client-app/models"
	"io"
	"log"
	"net/http"
	"time"
)

type IAuthService interface {
	AuthenticateGrant(authenticateGrantReq models.AuthAuthenticateGrantRequestModel) (models.AuthAuthenticateGrantResponseModel, error)
	GetOpenId(openIdReq models.AuthOpenIdRequestModel) (models.AuthOpenIdResponseModel, error)
}

type AuthService struct {
}

func (a *AuthService) AuthenticateGrant(authenticateGrantReq models.AuthAuthenticateGrantRequestModel) (models.AuthAuthenticateGrantResponseModel, error) {
	log.Println("Start AuthenticateGrant")
	serialized, err := json.Marshal(authenticateGrantReq)
	if err != nil {
		log.Println("error when marshalling authenticateGrantReq")
		return *new(models.AuthAuthenticateGrantResponseModel), errors.New("error when marshalling authenticateGrantReq")
	}
	bodyReader := bytes.NewReader(serialized)
	authGrantUrl := "http://localhost:1323/oauth2/authenticate-grant"
	req, err := http.NewRequest(http.MethodPost, authGrantUrl, bodyReader)
	if err != nil {
		log.Println("error when creating request")
		return *new(models.AuthAuthenticateGrantResponseModel), errors.New("error when creating request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client failed to do request")
		return *new(models.AuthAuthenticateGrantResponseModel), errors.New("error when creating request")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error readAll req body")
		return *new(models.AuthAuthenticateGrantResponseModel), errors.New("error when creating request")
	}

	result := models.BaseResponseModel[models.AuthAuthenticateGrantResponseModel]{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("error when unmarshalling data")
		return *new(models.AuthAuthenticateGrantResponseModel), errors.New("error when unmarshalling data")
	}

	log.Println("Finished marshalling authenticateGrantReq")
	return result.Data, nil
}

func (a *AuthService) GetOpenId(openIdReq models.AuthOpenIdRequestModel) (models.AuthOpenIdResponseModel, error) {
	log.Println("Start GetOpenId")
	getOpenIdUrl := "http://localhost:1323/player/get-my-profile"
	req, err := http.NewRequest(http.MethodGet, getOpenIdUrl, nil)
	if err != nil {
		log.Println("error when creating request")
		return *new(models.AuthOpenIdResponseModel), errors.New("error when creating request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openIdReq.Token))

	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client failed to do request")
		return *new(models.AuthOpenIdResponseModel), errors.New("error when exec request")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error readAll req body")
		return *new(models.AuthOpenIdResponseModel), errors.New("error when reading response body")
	}

	var result models.BaseResponseModel[models.AuthOpenIdResponseModel]
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("error when unmarshalling data")
		return *new(models.AuthOpenIdResponseModel), errors.New("error when unmarshal response body")
	}

	log.Println("finish GetOpenId")
	return result.Data, nil
}

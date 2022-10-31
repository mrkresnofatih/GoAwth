package services

import (
	"errors"
	"fmt"
	"github.com/mrkresnofatih/go-awth/models"
	"github.com/mrkresnofatih/go-awth/tools/jwt"
	"log"
	"strings"
	"time"
)

type IOauthService interface {
	GetConsent(getConsent models.OauthGetConsentRequestModel) (models.OauthGetConsentResponseModel, error)
	AgreeConsent(agreeRequest models.OauthAgreeConsentRequestModel) (models.OauthAgreeConsentResponseModel, error)
	RejectConsent(rejectRequest models.OauthRejectConsentRequestModel) (models.OauthRejectConsentResponseModel, error)
	AuthenticateGrant(authGrant models.OauthAuthenticateGrantRequestModel) (models.OauthAuthenticateGrantResponseModel, error)
}

type OauthService struct {
	PlayerService                    IPlayerService
	DeveloperApplicationService      IDeveloperApplicationService
	DeveloperApplicationGrantService IDeveloperApplicationGrantService
}

func (o *OauthService) AuthenticateGrant(authGrant models.OauthAuthenticateGrantRequestModel) (models.OauthAuthenticateGrantResponseModel, error) {
	grant, err := o.DeveloperApplicationGrantService.Get(models.DeveloperAppGrantGetRequestModel{
		DeveloperAppGrantId: authGrant.GrantId,
	})
	if err != nil {
		log.Println("grant is not found")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("grant not found")
	}

	grantedApp, err := o.DeveloperApplicationService.Read(models.DeveloperApplicationReadRequestModel{
		DeveloperApplicationId: grant.ApplicationId,
	})
	if err != nil {
		log.Println("granted app not found")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("granted app not found")
	}

	if grantedApp.DeveloperApplicationId != authGrant.ApplicationId {
		log.Println("grantedAppId incorrect")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("granted app credentials incorrect")
	}

	if grantedApp.Secret != authGrant.ApplicationSecret {
		log.Println("grantedAppCredentials incorrect")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("granted app credentials incorrect")
	}

	grantExpireTime, err := time.Parse(time.RFC3339, grant.ExpiresAt)
	isExpired := time.Now().After(grantExpireTime)
	if isExpired {
		log.Println("grant is expired")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("grant is expired")
	}

	basicTokenBuilder := &jwt.BasicJwtTokenBuilder{
		ExpiresAfter: time.Hour * 1,
	}
	withUsernameTokenBuilder := &jwt.UsernameJwtTokenBuilder{
		JwtTokenBuilder: basicTokenBuilder,
		Username:        grant.PlayerUsername,
	}
	withGrantScopesTokenBuilder := &jwt.ApplicationJwtTokenBuilder{
		JwtTokenBuilder: withUsernameTokenBuilder,
		GrantId:         grant.DeveloperAppGrantIdId,
		GrantScopes:     grant.Scope,
	}
	token, err := withGrantScopesTokenBuilder.Build()
	if err != nil {
		log.Println("failed to build granted app token")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("granted app token creation failed")
	}

	return models.OauthAuthenticateGrantResponseModel{
		GrantId:         grant.DeveloperAppGrantIdId,
		ExpiresAt:       grant.ExpiresAt,
		GrantToken:      token,
		PermittedScopes: grant.Scope,
	}, nil
}

func (o *OauthService) GetConsent(getConsent models.OauthGetConsentRequestModel) (models.OauthGetConsentResponseModel, error) {
	player, err := o.PlayerService.Get(models.PlayerGetRequestModel{
		Username: getConsent.PlayerUsername,
	})
	if err != nil {
		log.Println("failed to get player")
		return *new(models.OauthGetConsentResponseModel), errors.New("failed to get player")
	}

	scopeDefinitionMap := getScopeDefinitionsMap(getConsent.Scope)

	developerApp, err := o.DeveloperApplicationService.Read(models.DeveloperApplicationReadRequestModel{
		DeveloperApplicationId: getConsent.DeveloperApplicationId,
	})
	if err != nil {
		log.Println("failed to get developer app")
		return *new(models.OauthGetConsentResponseModel), errors.New("failed to get developer app")
	}

	log.Println("successfully get consent")
	return models.OauthGetConsentResponseModel{
		DeveloperApplicationId:       developerApp.DeveloperApplicationId,
		DeveloperApplicationImageUrl: developerApp.LogoUrl,
		DeveloperName:                developerApp.DeveloperName,
		DeveloperApplicationName:     developerApp.Name,
		PlayerUsername:               player.Username,
		PlayerImageUrl:               player.ImageUrl,
		ScopeDefinitions:             scopeDefinitionMap,
	}, nil
}

func (o *OauthService) AgreeConsent(agreeRequest models.OauthAgreeConsentRequestModel) (models.OauthAgreeConsentResponseModel, error) {
	app, err := o.DeveloperApplicationService.Read(models.DeveloperApplicationReadRequestModel{
		DeveloperApplicationId: agreeRequest.DeveloperApplicationId,
	})
	if err != nil {
		log.Println("app not found")
		return *new(models.OauthAgreeConsentResponseModel), errors.New("app not found")
	}

	newAppGrant, err := o.DeveloperApplicationGrantService.Create(models.DeveloperAppGrantCreateRequestModel{
		PlayerUsername: agreeRequest.PlayerUsername,
		Scope:          filterKnownScopes(agreeRequest.Scope),
		ExpiresAt:      time.Now().Add(time.Minute * 5).Format(time.RFC3339),
		ApplicationId:  agreeRequest.DeveloperApplicationId,
	})
	if err != nil {
		log.Println("error creating grant")
		return *new(models.OauthAgreeConsentResponseModel), errors.New("error creating grant")
	}

	log.Println("successfully created new app grant with grant-id: " + newAppGrant.DeveloperAppGrantId)
	return models.OauthAgreeConsentResponseModel{
		GrantId:     newAppGrant.DeveloperAppGrantId,
		RedirectUri: buildSuccessUri(app.SuccessRedirectUri, newAppGrant.DeveloperAppGrantId),
	}, nil
}

func (o *OauthService) RejectConsent(rejectRequest models.OauthRejectConsentRequestModel) (models.OauthRejectConsentResponseModel, error) {
	app, err := o.DeveloperApplicationService.Read(models.DeveloperApplicationReadRequestModel{
		DeveloperApplicationId: rejectRequest.DeveloperApplicationId,
	})
	if err != nil {
		log.Println("app not found")
		return *new(models.OauthRejectConsentResponseModel), errors.New("app not found")
	}

	log.Println("successfully rejected app grant")
	return models.OauthRejectConsentResponseModel{
		RedirectUri: app.FailedRedirectUri,
	}, nil
}

const (
	ScopeName_OpenID       = "openid"
	ScopeDefinition_OpenID = "Get Your User Profile"
)

func getScopeDefinitionsMap(scope string) map[string]string {
	scopeSplits := strings.Fields(scope)
	scopeMap := map[string]string{}
	for _, scopeName := range scopeSplits {
		switch scopeName {
		case ScopeName_OpenID:
			if _, ok := scopeMap["foo"]; !ok {
				scopeMap[scopeName] = ScopeDefinition_OpenID
			}
		}
	}
	return scopeMap
}

func filterKnownScopes(scope string) string {
	scopeSplits := strings.Fields(scope)
	scopeMap := map[string]string{}
	scopeString := ""
	for _, scopeName := range scopeSplits {
		switch scopeName {
		case ScopeName_OpenID:
			if _, ok := scopeMap["foo"]; !ok {
				scopeMap[scopeName] = ScopeDefinition_OpenID
				if len(scopeString) == 0 {
					scopeString = scopeName
				} else {
					scopeString = fmt.Sprintf("%s %s", scopeString, scopeName)
				}
			}
		}
	}
	return scopeString
}

func buildSuccessUri(redirectUri, grantId string) string {
	return fmt.Sprintf("%s?grantId=%s", redirectUri, grantId)
}

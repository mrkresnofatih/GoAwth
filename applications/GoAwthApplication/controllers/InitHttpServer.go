package controllers

import (
	"github.com/mrkresnofatih/go-awth/applications"
	"github.com/mrkresnofatih/go-awth/services"
	"sync"
)

func InitHttpServer(appRunState *sync.WaitGroup) {
	go func() {
		httpServerObj := &ApplicationServer{}

		playerController := &PlayerController{
			PlayerService: &services.PlayerService{
				GormClient: applications.GetGormMySqlInstance(),
			},
		}
		httpServerObj.AddController(playerController)

		developerController := &DeveloperController{
			DeveloperService: &services.DeveloperService{
				GormClient: applications.GetGormMySqlInstance(),
			},
		}
		httpServerObj.AddController(developerController)

		devAppsController := &DeveloperApplicationController{
			DeveloperApplicationService: &services.DeveloperApplicationService{
				GormClient: applications.GetGormMySqlInstance(),
			},
		}
		httpServerObj.AddController(devAppsController)

		httpServerObj.Initialize()
		httpServerObj.Router.Logger.Fatal(httpServerObj.Router.Start(":1323"))
		appRunState.Done()
	}()
}

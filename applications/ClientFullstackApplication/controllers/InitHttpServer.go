package controllers

import (
	"github.com/mrkresnofatih/go-awth-client-app/services"
	"sync"
)

func InitHttpServer(appRunState *sync.WaitGroup) {
	go func() {
		httpServerObj := &ApplicationServer{}

		authController := &AuthController{
			AuthService: &services.AuthService{},
		}
		httpServerObj.AddController(authController)

		httpServerObj.Initialize()
		httpServerObj.Router.Logger.Fatal(httpServerObj.Router.Start(":1324"))
		appRunState.Done()
	}()
}

package main

import (
	"github.com/mrkresnofatih/go-awth/applications"
	"github.com/mrkresnofatih/go-awth/controllers"
	"sync"
)

func main() {
	var appRunState sync.WaitGroup
	appRunState.Add(1)
	applications.RunGormMigration()
	controllers.InitHttpServer(&appRunState)
	appRunState.Wait()
}

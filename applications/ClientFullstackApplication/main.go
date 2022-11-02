package main

import (
	"github.com/mrkresnofatih/go-awth-client-app/controllers"
	"sync"
)

func main() {
	var appRunState sync.WaitGroup
	appRunState.Add(1)
	controllers.InitHttpServer(&appRunState)
	appRunState.Wait()
}

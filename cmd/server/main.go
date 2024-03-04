package main

import (
	"github.com/nickWoott/my-hex-api/internal/adapters/primary/http"
	database "github.com/nickWoott/my-hex-api/internal/adapters/secondary"
	testServices "github.com/nickWoott/my-hex-api/internal/core/services"
)

func main() {
	DatabaseConnection := database.NewLocalDatabase()
	testService := testServices.NewTestService(&DatabaseConnection)

	http := http.NewHttpServer()

	http.Run(testService)
}

//what do we need to do here?
// probably create a new server

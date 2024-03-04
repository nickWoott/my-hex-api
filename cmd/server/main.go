package main

import (
	"github.com/nickWoott/my-hex-api/internal/adapters/primary/http"
	testServices "github.com/nickWoott/my-hex-api/internal/core/services"
)

func main() {

	testService := testServices.NewTestService()

	http := http.NewHttpServer()

	http.Run(testService)
}

//what do we need to do here?
// probably create a new server

package main

import "github.com/nickWoott/my-hex-api/internal/adapters/primary/http"

func main() {

	http := http.NewHttpServer()

	http.Run()
}

//what do we need to do here?
// probably create a new server

package http

import (
	"github.com/gin-gonic/gin"
	testController "github.com/nickWoott/my-hex-api/internal/adapters/primary/http/controllers"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Run() {

	testController := testController.NewTestController()
	r := gin.Default()

	r.GET("/test", testController.SendTest)
	// this is where we create controllers
	// we can also use middleware here
	// here we will also set up our routes
	r.Run("localhost:3000")
}

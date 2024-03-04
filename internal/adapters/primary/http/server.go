package http

import (
	"github.com/gin-gonic/gin"
	testController "github.com/nickWoott/my-hex-api/internal/adapters/primary/http/controllers"
	testServices "github.com/nickWoott/my-hex-api/internal/core/services"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Run(service testServices.TestService) {

	testController := testController.NewTestController(service)
	r := gin.Default()

	r.GET("/test", testController.SendTest)
	r.Run("localhost:3000")
}

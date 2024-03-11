package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nickWoott/my-hex-api/internal/adapters/primary/http/controllers"
	"github.com/nickWoott/my-hex-api/internal/core/services"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Run(service services.StoryPointService) {

	postStoryPointController := controllers.NewPostStoryPointController(service)

	getStoryPointController := controllers.NewGetStoryPointController(service)
	r := gin.Default()

	r.GET("/storypoint", getStoryPointController.SendTest)
	r.POST("/storypoint",postStoryPointController.CreateStoryPoints )
	r.Run("localhost:3000")


}

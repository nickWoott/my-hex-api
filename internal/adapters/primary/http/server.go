package http

import (
	"github.com/gin-contrib/cors"
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

	deleteAllStoryPointsController := controllers.NewDeleteAllStoryPointsController(service)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/storypoint", getStoryPointController.SendTest)
	r.POST("/storypoint",postStoryPointController.CreateStoryPoints )
	r.DELETE("storypoints", deleteAllStoryPointsController.DeleteAllStoryPoints)
	r.Run("localhost:3001")


}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nickWoott/my-hex-api/internal/core/domain"
	"github.com/nickWoott/my-hex-api/internal/core/services"
)

type PostStoryPointsController struct {
	service services.StoryPointService
}

func NewPostStoryPointController(service services.StoryPointService) PostStoryPointsController {
	return PostStoryPointsController{
		service: service,
	}
}

func  (pspc *PostStoryPointsController) CreateStoryPoints(c *gin.Context) {

	
	var si domain.StorypointRequest

	if err := c.BindJSON(&si); err != nil {
		println("cannot bind json to storypoint struct")
		return
	}


	if err := pspc.service.SendStoryPoints(&si); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"convertedLead": si})
}

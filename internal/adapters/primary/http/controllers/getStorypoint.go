package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nickWoott/my-hex-api/internal/core/services"
)

type GetStoryPointController struct {
	service services.StoryPointService
}

func NewGetStoryPointController(service services.StoryPointService) GetStoryPointController {
	return GetStoryPointController{
		service: service,
	}
}

func (gspc *GetStoryPointController) SendTest(c *gin.Context) {

	id := c.Query("id")
	storyPoint, err := gspc.service.GetData(id)
	if err != nil {
	fmt.Println("there has been an error in the controller")
	}

	c.JSON(200, storyPoint)
}

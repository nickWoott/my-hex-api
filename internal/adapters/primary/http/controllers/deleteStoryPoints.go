package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nickWoott/my-hex-api/internal/core/services"
)

type DeleteAllStoryPointsController struct {
	service services.StoryPointService
}

func NewDeleteAllStoryPointsController(service services.StoryPointService) DeleteAllStoryPointsController {
return DeleteAllStoryPointsController{
	service: service,
}

}

func(dspc *DeleteAllStoryPointsController) DeleteAllStoryPoints(c *gin.Context) {
 err := dspc.service.DeleteAll()

if err != nil {
	c.JSON(400, err)
}

c.JSON(200, "succesfully deleted all storypoints")
}



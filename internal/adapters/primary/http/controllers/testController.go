package testController

import (
	"fmt"

	"github.com/gin-gonic/gin"
	testServices "github.com/nickWoott/my-hex-api/internal/core/services"
)

type TestController struct {
	service testServices.TestService
}

func NewTestController(service testServices.TestService) TestController {
	return TestController{
		service: service,
	}
}

func (tc *TestController) SendTest(c *gin.Context) {

	id := c.Query("id")
	fmt.Println(id, "HERE I AM I AM THE ID")
	idFromService := tc.service.GetData(id)

	c.JSON(200, gin.H{"message": "I am working", "id": idFromService})
}

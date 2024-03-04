package testController

import "github.com/gin-gonic/gin"

type TestController struct {
}

func NewTestController() TestController {
	return TestController{}
}

func (tc *TestController) SendTest(c *gin.Context) {
	c.JSON(200, gin.H{"message": "I am working"})
}

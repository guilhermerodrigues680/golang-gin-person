package apiv1

import "github.com/gin-gonic/gin"

type otherController struct{}

func NewOtherController() *otherController {
	return &otherController{}
}

func (pc *otherController) HandlePing(c *gin.Context) {
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

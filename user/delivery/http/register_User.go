package http

import (
	"chat/user"
	"github.com/gin-gonic/gin"
)

func RegisterHttpEndPointsUser(c *gin.Engine, us user.UseCaseUsers) {
	h := NewHandlerUser(us)

	user := c.Group("/users")
	{
		user.POST("/add", h.CreateUser)
	}
}

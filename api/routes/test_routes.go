package routes

import (
	"clean-architecture/api/responses"
	"clean-architecture/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestRoutes struct {
	handler infrastructure.Router
}

func NewTestRoutes(
	handler infrastructure.Router,
) *TestRoutes {
	return &TestRoutes{
		handler: handler,
	}
}

func (t TestRoutes) Setup() {
	t.handler.GET("/test", func(c *gin.Context) {
		responses.JSON(c, http.StatusOK, "Test message.")
	})
}

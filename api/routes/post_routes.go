package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
)

type PostRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	postController controllers.PostController
}

func NewPostRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	postController controllers.PostController,
) PostRoutes {
	return PostRoutes{
		logger:         logger,
		handler:        handler,
		postController: postController,
	}
}

func (pr PostRoutes) Setup() {
	pr.logger.Info("Setting post routes")
	api := pr.handler.Group("/")
	{
		api.GET("/posts/:id", pr.postController.GetOnePost())

		api.POST("/posts", pr.postController.CreatePost())
	}
}

package controllers

import (
	"clean-architecture/api/responses"
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/repository"
	"clean-architecture/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service    services.PostService
	repository repository.PostRepository
	logger     lib.Logger
}

func NewPostController(
	postService services.PostService,
	postRepository repository.PostRepository,
	logger lib.Logger,
) PostController {
	return PostController{
		service:    postService,
		repository: postRepository,
		logger:     logger,
	}
}

func (pc PostController) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		pc.logger.Info("PostController: Creating Post")

		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			pc.logger.Error(err)
			responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		post, err := pc.repository.CreatePost(post)
		if err != nil {
			pc.logger.Error(err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		responses.JSON(c, http.StatusOK, post)
	}
}

func (pc PostController) GetOnePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		pc.logger.Info("PostController: Retrieving Post by Id")

		postID := c.Param("id")

		post, err := pc.repository.GetOnePost(postID)
		if err != nil {
			pc.logger.Error(err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		responses.JSON(c, http.StatusOK, post)
	}
}

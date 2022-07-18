package repository

import (
	"clean-architecture/infrastructure"
	"clean-architecture/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{db}
}

func (r PostRepository) WithTrx(trxHandle *gorm.DB) PostRepository {
	r.Database.DB = trxHandle
	return r
}

func (pr PostRepository) GetOnePost(postID string) (post models.Post, err error) {
	return post, pr.First(&post, "id=?", postID).Error
}

func (pr PostRepository) CreatePost(post models.Post) (savedPost models.Post, err error) {
	return post, pr.Save(&post).Error
}

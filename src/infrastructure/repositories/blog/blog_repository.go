package blog

import (
	blog "akastra-mobile-api/src/infrastructure/database/models/blog"

	"gorm.io/gorm"
)


type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

func (r *blogRepository) GetArticles(limit, offset int) ([]blog.Article, int64, error) {
	var articles []blog.Article
	var total int64

	r.db.Model(&blog.Article{}).Count(&total)
	if err := r.db.Preload("Writer").Preload("Category").Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}
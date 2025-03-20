package blog

import (
	blog "akastra-mobile-api/src/infrastructure/database/models/blog"
)

type BlogRepository interface {
	GetArticles(limit, offset int) ([]blog.Article, int64, error)
}
package blog

import mBlog "akastra-mobile-api/src/infrastructure/database/models/blog"

type BlogUsecase interface {
	GetArticles(limit, offset int) ([]mBlog.Article, int64, error)
}
package blog

import (
	mBlog "akastra-mobile-api/src/infrastructure/database/models/blog"
	rBlog "akastra-mobile-api/src/infrastructure/repositories/blog"
)

type blogUsecase struct {
	blogRepo rBlog.BlogRepository
}

func NewBlogUsecase(blogRepo rBlog.BlogRepository) BlogUsecase {
	return &blogUsecase{blogRepo: blogRepo}
}

func (r *blogUsecase) GetArticles(limit, offset int) ([]mBlog.Article, int64, error) {
	return r.blogRepo.GetArticles(limit, offset)
}

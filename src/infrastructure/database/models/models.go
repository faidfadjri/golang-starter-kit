package models

import (
	"akastra-mobile-api/src/infrastructure/database/models/blogs"
	"akastra-mobile-api/src/infrastructure/database/models/users"
)

type (
	Article         = blogs.Article
	ArticleCategory = blogs.ArticleCategory
	User            = users.User
	UserRole        = users.UserRole
)
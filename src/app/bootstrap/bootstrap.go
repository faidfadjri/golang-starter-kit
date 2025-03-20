package bootstrap

import (
	uAuth "akastra-mobile-api/src/app/usecase/auth"
	uBlog "akastra-mobile-api/src/app/usecase/blog"
	"akastra-mobile-api/src/infrastructure/database"
	rAuth "akastra-mobile-api/src/infrastructure/repositories/auth"
	rBlog "akastra-mobile-api/src/infrastructure/repositories/blog"
	"akastra-mobile-api/src/interface/handler"
	"log"
)

type Dependencies struct {
	AuthHandler *handler.AuthHandler
	BlogHandler *handler.BlogHandler
}

func InitDependencies() *Dependencies {
	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	authRepo := rAuth.NewAuthRepository(db)
	authUsecase := uAuth.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	blogRepo := rBlog.NewBlogRepository(db)
	blogUsecase := uBlog.NewBlogUsecase(blogRepo)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	return &Dependencies{
		AuthHandler: authHandler,
		BlogHandler: blogHandler,
	}
}

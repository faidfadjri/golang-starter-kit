package handler

import (
	uBlog "akastra-mobile-api/src/app/usecase/blog"
	"akastra-mobile-api/src/interface/response"
	"log"
	"net/http"
)

type BlogHandler struct {
	blogUseCase uBlog.BlogUsecase
}

func NewBlogHandler(blogUsecase uBlog.BlogUsecase) *BlogHandler {
	return &BlogHandler{
		blogUseCase: blogUsecase,
	}
}

func (h *BlogHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred:", r)
			response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Internal server error", nil))
		}
	}()

	limit := 4
	offset := 2

	articles, total, err := h.blogUseCase.GetArticles(limit, offset)
	if err != nil {
		log.Printf("ERROR: Failed to get articles data: %v\n", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Server Error", nil))
		return
	}

	data := map[string]interface{}{
		"total":    total,
		"articles": articles,
	}

	response.JSONResponse(w, http.StatusOK, response.NewSuccessResponse("Articles fetched successfully", data))
}
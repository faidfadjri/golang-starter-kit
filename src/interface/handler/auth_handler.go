package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/interface/response"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserRegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err))
		return
	}

	if payload.Fullname == "" || payload.Email == "" || payload.Password == "" {
		log.Println("ERROR: Missing required fields")
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("All fields are required", nil))
		return
	}

	registeredUser, err := h.authUsecase.Register(payload)
	if err != nil {
		log.Printf("ERROR: Failed to register user: %v\n", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Failed to register user", err))
		return
	}

	responseData := struct {
		Fullname string  `json:"fullname"`
		Username *string `json:"username,omitempty"`
		Email    string  `json:"email"`
		Phone    string  `json:"phone"`
		Address  string  `json:"address"`
		Avatar   string  `json:"avatar"`
	}{
		Fullname: registeredUser.Fullname,
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Phone:    registeredUser.Phone,
		Address:  registeredUser.Address,
		Avatar:   registeredUser.Avatar,
	}

	response.JSONResponse(w, http.StatusCreated, response.NewSuccessResponse("User registered successfully", responseData))
}


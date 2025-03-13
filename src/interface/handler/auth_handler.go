package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/interface/response"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	validate *validator.Validate
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase, validate: validator.New()}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// var payload entities.UserLoginPayload
	
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserRegisterPayload

	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred:", r)
			response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Internal server error", nil))
		}
	}()

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err))
		return
	}

	if err := h.validate.Struct(payload); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			firstError := validationErrs[0]
			errorMessage := fmt.Sprintf("Field '%s':%s", firstError.Field(), firstError.Tag())
	
			log.Println("ERROR: Validation failed:", errorMessage)
			response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse(errorMessage, nil))
			return
		}
	
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err))
		return
	}
	

	registeredUser, err := h.authUsecase.Register(payload)
	if err != nil {
		log.Printf("ERROR: Failed to register user: %v\n", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Failed to register user", err))
		return
	}

	responseData := entities.UserRegisterPayload{
		Fullname: registeredUser.Fullname,
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Phone:    registeredUser.Phone,
		Address:  registeredUser.Address,
	}

	response.JSONResponse(w, http.StatusCreated, response.NewSuccessResponse("User registered successfully", responseData))
}


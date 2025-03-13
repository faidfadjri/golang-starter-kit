package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// var payload entities.UserLoginPayload
	
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload entities.UserRegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err))
		return
	}

	err := h.validate.Struct(payload)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			log.Println("ERROR: Invalid validation error:", err)
			response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Internal server error", nil))
			return
		}
	
		var validationErrors []string
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				validationErrors = append(validationErrors, fmt.Sprintf("Field '%s' failed validation: %s", e.Field(), e.Tag()))
			}
	
			log.Println("ERROR: Validation failed:", validationErrors)
			response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Validation failed", errors.New(strings.Join(validationErrors, "; "))))
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


package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"akastra-mobile-api/src/app/entities"
	auth "akastra-mobile-api/src/app/usecase/auth"
	jwtutil "akastra-mobile-api/src/infrastructure/jwt"
	"akastra-mobile-api/src/interface/response"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUsecase auth.AuthUsecase
	validate *validator.Validate
}

func NewAuthHandler(authUsecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase, validate: validator.New()}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred:", r)
			response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Internal server error", nil))
		}
	}()

	var payload entities.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", nil))
		return
	}

	if err := h.validate.Struct(payload); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			firstError := validationErrs[0]
			errorMessage := fmt.Sprintf("Field '%s': %s", firstError.Field(), firstError.Tag())

			log.Println("ERROR: Validation failed:", errorMessage)
			response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse(errorMessage, nil))
			return
		}

		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", nil))
		return
	}

	loginUser, err := h.authUsecase.Login(payload)
	if err != nil {
		log.Printf("ERROR: Failed to login user: %v\n", err)
		response.JSONResponse(w, http.StatusUnauthorized, response.NewErrorResponse("Invalid email or password", nil))
		return
	}

	token, err := jwtutil.GenerateToken(jwtutil.Claims{
		UserID:   loginUser.ID,
		Fullname: loginUser.Fullname,
		Email:    loginUser.Email,
		Role:     loginUser.Role.Name,
		Phone:    loginUser.Phone,
		Address:  loginUser.Address,
		Avatar:   loginUser.Avatar,
	}, 24*time.Hour)

	if err != nil {
		log.Println("ERROR: Failed to generate token:", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Failed to generate token", nil))
		return
	}

	response.JSONResponse(w, http.StatusOK, response.NewSuccessResponse("Login succeeded", map[string]string{
		"token": token,
	}))
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


package handler

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/interface/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		log.Println("ERROR: Failed to get users:", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Failed to get users", err))
		return
	}

	response.JSONResponse(w, http.StatusOK, response.NewSuccessResponse("Users retrieved successfully", users))
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("ERROR: Invalid ID:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid ID", err))
		return
	}

	user, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		log.Println("ERROR: User not found:", err)
		response.JSONResponse(w, http.StatusNotFound, response.NewErrorResponse("User not found", err))
		return
	}

	response.JSONResponse(w, http.StatusOK, response.NewSuccessResponse("User retrieved successfully", user))
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	
	var payload entities.UserRegisterPayload

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("ERROR: Invalid request body:", err)
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("Invalid request body", err))
		return
	}

	// Inpout Validation
	if payload.Fullname == "" || payload.Email == "" || payload.Password == "" {
		log.Println("ERROR: Missing required fields")
		response.JSONResponse(w, http.StatusBadRequest, response.NewErrorResponse("All fields are required", nil))
		return
	}

	// Buat objek User
	user := entities.User{
		Name:     payload.Fullname,
		Email:    payload.Email,
		Password: payload.Password,
	}

	// Simpan ke database
	err = h.userUsecase.CreateUser(user)
	if err != nil {
		log.Printf("ERROR: Failed to create user: %v\n", err)
		response.JSONResponse(w, http.StatusInternalServerError, response.NewErrorResponse("Failed to create user", err))
		return
	}

	log.Printf("INFO: User created successfully - Name: %s, Email: %s\n", user.Name, user.Email)
	response.JSONResponse(w, http.StatusCreated, response.NewSuccessResponse("User created successfully", user))
}

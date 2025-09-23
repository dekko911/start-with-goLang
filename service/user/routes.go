package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dekko911/start-with-goLang/config"
	"github.com/dekko911/start-with-goLang/service/auth"
	"github.com/dekko911/start-with-goLang/types"
	"github.com/dekko911/start-with-goLang/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// ini adalah tempat menaruh routes + business logic.

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	r.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)

	// admin routes
	r.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON payload request
	payload := types.LoginUserPayload{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// checking the email
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// comparing the password payload req and the actual password
	if !auth.CompareHashedPassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// create token
	secret := []byte(config.Env.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"token":  token,
	})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// check if user has exists at table users
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hash the password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if doesn't, create a new one
	err = h.store.CreateUser(types.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	user := map[string]string{
		"Name":     payload.Name,
		"Email":    payload.Email,
		"Username": payload.Username,
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":  http.StatusCreated,
		"message": "User Created!",
		"user":    user,
	})
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid type user ID"))
		return
	}

	user, err := h.store.GetUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

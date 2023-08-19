package handler

import (
	"devstream-rest-api/internal/model"
	"devstream-rest-api/internal/repository"
	"devstream-rest-api/internal/util/appauth"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/apphttp"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) AuthHandler {
	return AuthHandler{
		userRepo: userRepo,
	}
}

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var params Login

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	user, err := h.userRepo.GetByUsername(params.Username)

	if err != nil {
		apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
		return
	}

	e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if e != nil {
		apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
		return
	}

	token, err := appauth.Generate(user.Id)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, map[string]string{
		"token": token,
	})
}

func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var params Register

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	user := &model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
		Username: params.Username,
	}

	err := h.userRepo.Create(user)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, "Registration successful", 201)
}

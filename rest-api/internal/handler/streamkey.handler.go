package handler

import (
	"devstream-rest-api/internal/repository"
	"devstream-rest-api/internal/util/appauth"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/apphttp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type StreamKeyHandler struct {
	streamKeyRepo repository.StreamKeyRepository
}

func NewStreamKeyHandler(streamKeyRepo repository.StreamKeyRepository) StreamKeyHandler {
	return StreamKeyHandler{
		streamKeyRepo: streamKeyRepo,
	}
}

type StreamKeyParams struct {
	Name string `json:"name" validate:"required"`
}

func (h StreamKeyHandler) GetStreamKeysForUser(w http.ResponseWriter, r *http.Request) {
	userId := appauth.GetUserIdFromContext(r)

	if userId == nil {
		apphttp.WriteJSONResponse(w, apperror.ServerError)
		return
	}

	keys, err := h.streamKeyRepo.GetByUserId(*userId)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, keys)
}

func (h StreamKeyHandler) CreateStreamKey(w http.ResponseWriter, r *http.Request) {
	userId := appauth.GetUserIdFromContext(r)

	if userId == nil {
		apphttp.WriteJSONResponse(w, apperror.ServerError)
		return
	}

	var params StreamKeyParams

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	keys, err := h.streamKeyRepo.Create(*userId, params.Name)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, keys, 201)
}

func (h StreamKeyHandler) UpdateStreamKey(w http.ResponseWriter, r *http.Request) {
	userId := appauth.GetUserIdFromContext(r)

	if userId == nil {
		apphttp.WriteJSONResponse(w, apperror.ServerError)
		return
	}

	var params StreamKeyParams

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	idStr := chi.URLParam(r, "id")

	id, e := strconv.ParseInt(idStr, 10, 64)

	if e != nil {
		apphttp.WriteJSONResponse(w, &apperror.BadRequest)
		return
	}

	keys, err := h.streamKeyRepo.Update(id, *userId, params.Name)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, keys, 201)
}

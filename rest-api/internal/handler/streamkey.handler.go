package handler

import (
	"devstream-rest-api/internal/repository"
	"devstream-rest-api/internal/util/appauth"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/apphttp"
	"net/http"
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

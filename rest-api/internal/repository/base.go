package repository

import (
	"devstream-rest-api/internal/model"
	"devstream-rest-api/internal/util/apperror"
)

type UserRepository interface {
	Create(*model.User) *apperror.AppError
	GetByEmail(string) (*model.User, *apperror.AppError)
	GetByUsername(string) (*model.User, *apperror.AppError)
}

type StreamKeyRepository interface {
	GetByUserId(int64) (*[]model.StreamKey, *apperror.AppError)
	Create(int64, string) (*model.StreamKey, *apperror.AppError)
	Update(int64, int64, string) (*model.StreamKey, *apperror.AppError)
}

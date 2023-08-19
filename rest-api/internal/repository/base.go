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

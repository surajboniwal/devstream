package apperror

import (
	"devstream-rest-api/internal/util/applogger"

	"github.com/lib/pq"
)

var logger = applogger.New("apperror")

type AppError struct {
	Tag         string `json:"tag"`
	UserMessage string `json:"message"`
	Code        int    `json:"-"`
}

var UnauthorizedError AppError = AppError{
	Tag:         "global",
	UserMessage: "Unauthorized",
	Code:        401,
}

var ServerError AppError = AppError{
	Tag:         "global",
	UserMessage: "Something went wrong",
	Code:        500,
}

func Parse(err error) *AppError {
	logger.E(err)
	if e, ok := err.(*pq.Error); ok {
		return parsePgError(e)
	}

	return &ServerError
}

func parsePgError(err *pq.Error) *AppError {
	return &AppError{
		Tag:         err.Column,
		UserMessage: DBErrorMap[err.Constraint],
		Code:        400,
	}
}

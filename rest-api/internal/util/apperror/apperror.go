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

var BadRequest AppError = AppError{
	Tag:         "global",
	UserMessage: "Bad request",
	Code:        400,
}

var NotFoundError AppError = AppError{
	Tag:         "global",
	UserMessage: "Not found",
	Code:        404,
}

func Parse(err error) *AppError {

	logger.E(err)

	if e, ok := err.(*pq.Error); ok {
		return parsePgError(e)
	}

	if err.Error() == "sql: no rows in result set" {
		return &NotFoundError
	}

	return &BadRequest
}

func parsePgError(err *pq.Error) *AppError {
	return &AppError{
		Tag:         err.Column,
		UserMessage: DBErrorMap[err.Constraint],
		Code:        400,
	}
}

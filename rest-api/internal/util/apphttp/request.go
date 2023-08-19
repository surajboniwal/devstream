package apphttp

import (
	"devstream-rest-api/internal/util/apperror"
	"encoding/json"
	"io"
	"net/http"
)

func ParseAndValidate(r *http.Request, data interface{}) *apperror.AppError {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return &apperror.AppError{
			Tag:         "global",
			UserMessage: "Error while parsing",
			Code:        400,
		}
	}

	if err := r.Body.Close(); err != nil {
		return &apperror.AppError{
			Tag:         "global",
			UserMessage: "Something went wrong",
			Code:        500,
		}
	}

	if err := json.Unmarshal(body, data); err != nil {
		return &apperror.AppError{
			Tag:         "global",
			UserMessage: "Error while parsing",
			Code:        400,
		}
	}

	if err := ValidateParam(data); err != nil {
		return err
	}

	return nil
}

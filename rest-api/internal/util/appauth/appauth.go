package appauth

import (
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/apphttp"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/o1egl/paseto"
)

var secretKey string

func Init(key string) {
	secretKey = key
}

func Generate(userid int64) (string, *apperror.AppError) {
	now := time.Now()
	exp := now.Add(time.Minute * 30)
	nbt := now
	t := paseto.NewV2()

	jsonToken := paseto.JSONToken{
		Issuer:     "devstream-rest-api",
		IssuedAt:   now,
		Subject:    fmt.Sprintf("%v", userid),
		Expiration: exp,
		NotBefore:  nbt,
	}

	token, err := t.Encrypt([]byte(secretKey), jsonToken, nil)

	if err != nil {
		return "", apperror.Parse(err)
	}

	return token, nil
}

func Validate(token string) (int64, *apperror.AppError) {
	var newJsonToken paseto.JSONToken
	t := paseto.NewV2()
	err := t.Decrypt(token, []byte(secretKey), &newJsonToken, nil)

	if err != nil {
		return 0, apperror.Parse(err)
	}

	if newJsonToken.Expiration.Before(time.Now()) {
		return 0, &apperror.UnauthorizedError
	}

	i, err := strconv.ParseInt(newJsonToken.Subject, 10, 64)
	if err != nil {
		return 0, apperror.Parse(err)
	}

	return i, nil
}

type ContextKey string

const userIdKey ContextKey = "user_id"

func GetUserIdFromContext(r *http.Request) *int64 {
	userId := r.Context().Value(userIdKey)

	if userId == nil {
		return nil
	}

	userIdInt64 := userId.(int64)

	return &userIdInt64
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := Validate(token)

		if err != nil {
			apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

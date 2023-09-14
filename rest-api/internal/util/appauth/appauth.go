package appauth

import (
	"context"
	"devstream-rest-api/internal/util/apperror"
	"devstream-rest-api/internal/util/apphttp"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func Init(key string) {
	secretKey = []byte(key)
}

type JWTClaim struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func Generate(userid int64) (string, *apperror.AppError) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", apperror.Parse(err)
	}

	return tokenString, nil
}

func Validate(tokenString string) *int64 {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		return nil
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil
	}

	return &claims.UserId
}

type ContextKey string

const userIdKey ContextKey = "user_id"

func GetUserIdFromContext(r *http.Request) *int64 {
	userId := r.Context().Value(userIdKey)

	if userId == nil {
		return nil
	}

	userIdInt64 := userId.(*int64)

	return userIdInt64
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId := Validate(token)

		if userId == nil {
			apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("auth_token")
		if err != nil {
			http.Error(writer, "unauthorized: missing auth token", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(writer, "unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			type AuthData struct {
				Email string
				Name  string
			}

			authData := AuthData{
				Email: fmt.Sprint(claims["email"]),
				Name:  fmt.Sprint(claims["name"]),
			}

			ctx := context.WithValue(request.Context(), authData, authData)
			request = request.WithContext(ctx)
		}

		next.ServeHTTP(writer, request)
	})
}

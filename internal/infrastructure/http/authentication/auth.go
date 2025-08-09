package authHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domainErrors "jobs.api.com/internal/domain/errors"
	authenticationUseCase "jobs.api.com/internal/usecases/authentication"
)

type AuthHandler struct {
	AuthUseCase authenticationUseCase.Authenticator
}

func NewAuthHandler(usecase authenticationUseCase.Authenticator) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: usecase,
	}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	type loginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginData loginDTO

	jsonErr := json.NewDecoder(request.Body).Decode(&loginData)

	if jsonErr != nil {
		http.Error(writer, "wrong payload", http.StatusBadRequest)
		return
	}
	fmt.Println("login data", loginData)

	user, err := handler.AuthUseCase.Login(loginData.Email, loginData.Password)

	if errors.Is(err, domainErrors.ErrUserNotFound) || errors.Is(err, domainErrors.ErrInvalidPasswordLogin) {
		http.Error(writer, "email or password invalid", http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(writer, "something went wrong", http.StatusInternalServerError)
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	cookieExpirationTime := time.Now().Add(48 * time.Hour)

	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"exp":   cookieExpirationTime.Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("token error", err, jwtSecret)
		http.Error(writer, "could not create token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  cookieExpirationTime,
	})

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("login successful"))

}

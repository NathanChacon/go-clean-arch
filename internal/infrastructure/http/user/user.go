package userHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	domainErrors "jobs.api.com/internal/domain/errors"
	userUseCase "jobs.api.com/internal/usecases/user"
)

type UserDto struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	UserUseCase userUseCase.UseCase
}

func NewUserHandler(usecase userUseCase.UseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: usecase,
	}
}

func (userHandler *UserHandler) GetUserById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if id == "" {
		http.Error(writer, "Missing user ID in URL", http.StatusBadRequest)
		return
	}

	user, err := userHandler.UserUseCase.GetById(id)

	if err != nil {
		fmt.Print("error", err)
		http.Error(writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	newUserDto := UserDto{
		Name:  user.Name,
		Email: user.Email,
		Bio:   user.Bio,
	}

	if err := json.NewEncoder(writer).Encode(newUserDto); err != nil {
		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
		return
	}

}

func (userHandler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var userData userUseCase.CreateAccountParams

	json.NewDecoder(request.Body).Decode(&userData)

	err := userHandler.UserUseCase.CreateUser(userData)

	if errors.Is(err, domainErrors.ErrInvalidUserCredential) {
		http.Error(writer, "Invalid credential", http.StatusBadRequest)
		return
	}

	if errors.Is(err, domainErrors.ErrInvalidEmailFormat) {
		http.Error(writer, "Invalid email", http.StatusBadRequest)
		return
	}

	if errors.Is(err, domainErrors.ErrUserAlreadyRegistered) {
		http.Error(writer, "User already registered", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

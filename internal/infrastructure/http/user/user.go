package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	userUseCase "jobs.api.com/internal/usecases/user"
)

type UserDto struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
}

type UserHandler struct {
	UserUseCase userUseCase.UseCase
}

func NewUserHandler(usecase userUseCase.UseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: usecase,
	}
}

// improve error validation
// isValidPayload ?
// did the user was found ?
// did was a internal error ?

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

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iggyster/lets-go-chat/internal/user"
)

type (
	Register struct {
		Repo user.UserRepo
	}
)

func (handler *Register) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := decodeAuthRequest(w, req)
	if err := handler.validate(&data); err.Count() != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Errors)

		return
	}

	usr := user.New(data.Username, data.Password)
	handler.Repo.Save(usr)

	json.NewEncoder(w).Encode(usr)
}

func (handler *Register) validate(data *AuthRequest) Errors {
	//TODO handler.Validator.validate(data, handler.getConstraints())
	errors := Errors{}
	if handler.Repo.IsExists(data.Username) {
		errors.AddError("userName", "User already exists", "Change the username")
	}
	if data.Username == "" {
		errors.AddError("userName", "User is empty", "Change the username")
	}
	if len(data.Username) < 3 {
		errors.AddError(
			"userName",
			"User is too short",
			"Username must be greater or equal to 3 characters",
		)
	}
	if data.Password == "" {
		errors.AddError("password", "Password is empty", "Change the username")
	}
	if len(data.Password) <= 8 || len(data.Password) > 32 {
		errors.AddError(
			"password",
			"Password is invalid",
			"Password must be more then 8 and less then 32 characters long",
		)
	}
	return errors
}

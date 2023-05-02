package handler

import (
	"encoding/json"
	"github.com/iggyster/lets-go-chat/internal/user"
	"net/http"
)

type RegisterData struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResource struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

func Register(resp http.ResponseWriter, req *http.Request) {
	data := RegisterData{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(resp, "Failed decoding", http.StatusInternalServerError)
		return
	}

	if errors := validate(&data); errors.Count() > 0 {
		SendErrors(resp, errors, http.StatusBadRequest)
		return
	}

	u := user.New(data.Username, data.Password)

	user.Repository.Save(u)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	err = json.NewEncoder(resp).Encode(RegisterResource{u.Id, u.Username})
	if err != nil {
		return
	}
}

func validate(data *RegisterData) Errors {
	errors := Errors{}
	if user.Repository.IsExists(data.Username) {
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

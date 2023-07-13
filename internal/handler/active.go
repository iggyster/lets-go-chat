package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iggyster/lets-go-chat/internal/user"
)

type Active struct {
	Repo user.UserRepo
}

//	@Summary		Active
//	@Description	Gets users list connected to the chat
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	user.User
//	@Router			/user/active [get]
func (a *Active) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := a.Repo.FindActivated()

	json.NewEncoder(w).Encode(users)
}

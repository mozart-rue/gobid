package api

import (
	"errors"
	"net/http"

	"github.com/mozart-rue/gobid/internal/jsonutils"
	"github.com/mozart-rue/gobid/internal/service"
	"github.com/mozart-rue/gobid/internal/usecase/user"
)

func (api *Api) handleSignUpUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodevalidJson[user.CreateUserRequest](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(
		r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)

	if err != nil {
		if errors.Is(err, service.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusConflict, map[string]any{
				"message": "username or email already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusConflict, map[string]any{
		"userId": id,
	})
}

func (api *Api) handleSignInUser(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodevalidJson[user.SignInUserRequest](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {}

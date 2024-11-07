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
	data, problems, err := jsonutils.DecodevalidJson[user.SignInUserRequest](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.SignInUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "invalid credentials",
			})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong while authenticating the user",
		})
		return
	}

	err = api.Session.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong while authenticating the user",
		})
		return
	}

	api.Session.Put(r.Context(), "AuthenticatedUserID", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "User authenticated",
	})
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Session.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong while logging out the user",
		})
		return
	}

	api.Session.Remove(r.Context(), "AuthenticatedUserID")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "User logged out",
	})
}

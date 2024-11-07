package api

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/mozart-rue/gobid/internal/jsonutils"
)

func (api *Api) HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"token": token,
	})
}

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		// If not, return an error
		// If the user is authenticated, call the next handler
		if !api.Session.Exists(r.Context(), "AuthenticatedUserID") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "unauthorized",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

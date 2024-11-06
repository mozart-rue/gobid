package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/mozart-rue/gobid/internal/service"
)

type Api struct {
	Router      *chi.Mux
	UserService service.UserService
	Session     *scs.SessionManager
}

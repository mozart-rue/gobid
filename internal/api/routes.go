package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {
	csrfMiddleware := csrf.Protect([]byte(os.Getenv("CSRF_KEY")))
	csrf.Secure(false) // Allow request from HTTP and HTTPS

	api.Router.Use(csrfMiddleware)

	api.Router.Use(middleware.RequestID)
	api.Router.Use(middleware.Logger)
	api.Router.Use(api.Session.LoadAndSave)
	api.Router.Use(middleware.Recoverer)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrftoken", api.HandleGetCSRFToken)
			r.Route("/users", func(r chi.Router) {
				r.Post("/sign-up", api.handleSignUpUser)
				r.Post("/sign-in", api.handleSignInUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogOutUser)
			})
		})
	})
}

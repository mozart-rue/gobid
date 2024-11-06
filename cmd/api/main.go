package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mozart-rue/gobid/internal/api"
	"github.com/mozart-rue/gobid/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("GOBID_DATABASE_USER"), os.Getenv("GOBID_DATABASE_PASSWORD"), os.Getenv("GOBID_DATABASE_HOST"), os.Getenv("GOBID_DATABASE_PORT"), os.Getenv("GOBID_DATABASE_NAME")))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *service.NewUserService(pool),
		Session:     s,
	}

	api.BindRoutes()

	fmt.Println("Server is running on port :3080")

	if err := http.ListenAndServe(":3080", api.Router); err != nil {
		panic(err)
	}
}

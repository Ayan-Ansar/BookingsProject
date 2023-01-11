package main

import (
	"github.com/Ayan-Ansar/bookings/pkg/config"
	"github.com/Ayan-Ansar/bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// middleware actually allows you to process a request as it comes into your web application and perform some action on it
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

// how to write middlewares

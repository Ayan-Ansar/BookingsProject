package main

import (
	"net/http"

	"github.com/Ayan-Ansar/bookings/internal/config"
	"github.com/Ayan-Ansar/bookings/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// middleware actually allows you to process a request as it comes into your web application and perform some action on it
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/suite", handlers.Repo.Suite)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/basic-room", handlers.Repo.BasicRoom)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	mux.Post("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}

// how to write middlewares

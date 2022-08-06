package main

import (
	"net/http"

	"github.com/Durotimicodes/bookings/pkg/config"
	"github.com/Durotimicodes/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	//GO-CHI
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	// mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//to handle static files we use fileserver
	fileSever := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileSever))
	return mux

}

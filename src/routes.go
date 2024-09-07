package main

import (
	"net/http"

	"github.com/VoltealProductions/TheArchiveReader/services/site"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerRouters() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	router.Get("/", site.GetHomePage)

	return router
}

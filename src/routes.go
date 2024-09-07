package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerRouters() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", getHomePage)
	router.Get("/about", getAboutPage)

	return router
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
}

func getAboutPage(w http.ResponseWriter, r *http.Request) {
}

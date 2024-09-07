package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerRouters() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return router
}

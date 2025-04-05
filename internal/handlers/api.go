package handlers

import (
	"github.com/SaiAungMinKhant/goAPI/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {

	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /accounts route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})


}
package account_service

import (
	"account-service/internal/ui/account-service/handlers"
	"account-service/internal/ui/account-service/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(usecases *ports.UseCases) http.Handler {
	/*mux := http.NewServeMux()
	mux.Handle("/api/login", handlers.NewConfirmAuthHandler(usecases.ConfirmAuth, "https://localhost:3000"))
	return mux*/

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/health-check", handlers.NewConfirmAuthHandler(usecases.ConfirmAuth, "https://localhost:3000").ServeHTTP)
	return router
}

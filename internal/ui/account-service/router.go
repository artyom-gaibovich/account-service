package account_service

import (
	"account-service/internal/ui/account-service/handlers"
	"account-service/internal/ui/account-service/ports"
	"net/http"
)

func NewRouter(usecases *ports.UseCases) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/login", handlers.NewConfirmAuthHandler(usecases.ConfirmAuth, "https://localhost:3000"))
	return mux

}

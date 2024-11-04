package handlers

import (
	"account-service/internal/ui/account-service/ports"
	"github.com/go-chi/render"
	"net/http"
)

type ConfirmAuthHandler struct {
	usecase ports.ConfirmAuthUseCase
	baseURL string
}

type Response struct {
	Status  string `json:"status"` // Error, ok
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewConfirmAuthHandler(usecase ports.ConfirmAuthUseCase, baseURL string) *ConfirmAuthHandler {
	return &ConfirmAuthHandler{
		usecase: usecase,
		baseURL: baseURL,
	}
}

func (h *ConfirmAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Message: "OK",
	})

}

package handlers

import (
	"account-service/internal/ui/account-service/ports"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
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
	var request ConfirmAuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.usecase.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	render.JSON(w, r, Response{
		Message: "OK",
	})

}

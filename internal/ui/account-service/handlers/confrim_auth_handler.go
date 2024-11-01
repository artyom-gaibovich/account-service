package handlers

import (
	"account-service/internal/ui/account-service/ports"
	"encoding/json"
	"fmt"
	"net/http"
)

type ConfirmAuthHandler struct {
	usecase ports.ConfirmAuthUseCase
	baseURL string
}

func NewConfirmAuthHandler(usecase ports.ConfirmAuthUseCase, baseURL string) *ConfirmAuthHandler {
	return &ConfirmAuthHandler{
		usecase: usecase,
		baseURL: baseURL,
	}
}

func (h *ConfirmAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Print("Это у нас GET запрос !!!")
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var requestBody struct {
			Username   string `json:"username"`
			Id         int    `json:"id"`
			TelegramId int    `json:"telegram_id"`
		}
		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		fmt.Printf("POST request: state=%s, code=%s\n", requestBody.Id, requestBody.Username, requestBody.TelegramId)
	}
	w.WriteHeader(http.StatusNotFound)

}

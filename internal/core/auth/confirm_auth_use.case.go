package auth

import (
	"account-service/internal/core/shared"
)

type ConfirmAuthRequest struct {
	ID    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}
type ConfirmAuthUseCase struct {
	userManager UserManager
	idGenerator shared.IDGenerator
}

func NewConfirmAuthUseCase(userManager UserManager, idGenerator shared.IDGenerator) *ConfirmAuthUseCase {
	return &ConfirmAuthUseCase{
		userManager: userManager,
		idGenerator: idGenerator,
	}
}

func (u *ConfirmAuthUseCase) Execute() error {
	return nil
}

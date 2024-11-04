package auth

import (
	"account-service/internal/core/shared"
	"context"
	"fmt"
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
func (u *ConfirmAuthUseCase) createUser(ctx context.Context, request handlers.ConfirmAuthRequest) (User, error) {
	user := User{
		ID:    u.idGenerator.Generate(),
		Email: request.Email,
		Name:  request.Name,
	}
	err := u.userManager.Create(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("error creatinng user on ConfirmOAuth2UseCase: %w", err)
	}
	return user, nil
	return User{}, nil
}

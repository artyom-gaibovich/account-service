package auth

import (
	"account-service/internal/core/shared"
	"context"
	"fmt"
)

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

func (u *ConfirmAuthUseCase) Execute(ctx context.Context, state, code string) (err error) {
	fmt.Println(state, code, code)
	fmt.Println(state, code, code)
	//return u.createUser(ctx)
	return nil
}
func (u *ConfirmAuthUseCase) createUser(ctx context.Context, providerUser ProviderUser) (User, error) {
	/*user := User{
		ID:             u.idGenerator.Generate(),
		Email:          providerUser.Email,
		Name:           providerUser.Name,
		ProviderUserID: providerUser.ID,
	}
	err := u.userManager.Create(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("error creatinng user on ConfirmOAuth2UseCase: %w", err)
	}
	return user, nil*/
	return User{}, nil
}

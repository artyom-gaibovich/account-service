package ports

import "context"

type UseCases struct {
	CreateUser  CreateUserUseCase
	DeleteUser  DeleteUserUseCase
	FindUser    FindUserUseCase
	ConfirmAuth ConfirmAuthUseCase
}

type CreateUserUseCase interface {
	Execute() (string, error)
}

type DeleteUserUseCase interface {
	Execute() (string, error)
}

type FindUserUseCase interface {
	Execute() (string, error)
}

type ConfirmAuthUseCase interface {
	Execute(ctx context.Context, state, code string) error
}
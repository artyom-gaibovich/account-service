package auth

import "context"

type UserManager interface {
	Create(ctx context.Context, user User) error
}

package postgres

import (
	"account-service/internal/core/auth"
	"account-service/pkg/dbrepository"
	"context"
	"database/sql"
	"fmt"
)

type UserManager struct {
	*dbrepository.Base
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Base: dbrepository.NewBase(db)}
}

func (r *UserManager) Create(ctx context.Context, user auth.User) error {
	err := r.Insert(ctx, "auth_users", map[string]interface{}{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"provider_user_id": user.ProviderUserID,
	})

	if err != nil {
		return fmt.Errorf("error on CreateUser: %w", err)
	}

	return nil
}

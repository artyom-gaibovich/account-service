package usermanager

import (
	"account-service/internal/adapters/usermanager/postgres"
	"database/sql"
)

func NewPostgresUserManager(db *sql.DB) *postgres.UserManager {
	return postgres.NewUserManager(db)
}

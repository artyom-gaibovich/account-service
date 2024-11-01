package app

import (
	"account-service/internal/core/auth"
	"account-service/pkg/env"
)

type Context struct {
	Env string

	Port int

	PostgresURL    string
	UserRepository auth.UserRepository
	UserManager    auth.UserManager
}

func NewContext() *Context {
	return &Context{
		Port: env.GetInt("PORT", 3000),
	}
}

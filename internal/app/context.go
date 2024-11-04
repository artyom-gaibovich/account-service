package app

import (
	"account-service/internal/adapters/idgenerator"
	"account-service/internal/adapters/usermanager"
	"account-service/internal/core/auth"
	"account-service/internal/core/shared"
	account_service "account-service/internal/ui/account-service"
	webports "account-service/internal/ui/account-service/ports"
	"account-service/pkg/env"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"
)

type Context struct {
	Env string

	Port int

	PostgresURL string

	db          *sql.DB
	userManager auth.UserManager
}

func NewContext() *Context {
	return &Context{
		Port:        env.GetInt("PORT", 3000),
		PostgresURL: env.GetString("POSTGRES_URL", "postgres://admin:root@localhost:5432/account_v1?sslmode=disable"),
	}
}

func (c *Context) UseCases() *webports.UseCases {
	return &webports.UseCases{
		ConfirmAuth: c.ConfirmAuthUseCase(),
	}
}

func (c *Context) Server() *account_service.Server {
	return account_service.NewServer(c.Port, c.Router(), c.Logger())
}

func (c *Context) Router() http.Handler {
	return account_service.NewRouter(c.UseCases())
}

func (c *Context) ConfirmAuthUseCase() *auth.ConfirmAuthUseCase {
	return auth.NewConfirmAuthUseCase(c.UserManager(), c.IDGenerator())
}

func (c *Context) IDGenerator() shared.IDGenerator {
	return idgenerator.NewUUIDGenerator()
}

func (c *Context) resolvePostgresURL() string {
	return c.PostgresURL
}

func (c *Context) DB() *sql.DB {
	if c.db == nil {
		url := c.resolvePostgresURL()
		fmt.Print(url)
		db, err := sql.Open("pgx", c.resolvePostgresURL())
		if err != nil {
			panic(err)
		}
		c.db = db
	}
	return c.db
}

//REPOSITORIES

func (c *Context) UserManager() auth.UserManager {
	if c.userManager == nil {
		c.userManager = usermanager.NewPostgresUserManager(c.DB())
	}
	return c.userManager
}

func (c *Context) Logger() *log.Logger {
	return log.New(os.Stdout, "web: ", log.Ldate|log.Ltime|log.LUTC)
}

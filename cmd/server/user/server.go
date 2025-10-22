package server

import (
	userService "auth-sso/internals/application/user"
	"auth-sso/internals/infrastructure/db"
	userhandler "auth-sso/internals/infrastructure/handler/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupUser(pool *pgxpool.Pool) *userhandler.Handler {
	repo := db.NewUserRepository(pool)
	service := userService.NewService(repo)
	handler := userhandler.NewHandler(service)
	return handler
}

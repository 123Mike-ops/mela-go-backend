package ports

import (
	"auth-sso/internals/domain/user"
	"context"
)

type Repository interface {
    GetByID(id int) (*user.User, error)
    Create(ctx context.Context,u *user.User)(*user.User, error)
    ExistsByEmailOrPhone(ctx context.Context, email, phoneNumber string) (bool, error)
}
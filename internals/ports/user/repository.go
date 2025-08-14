package ports

import "auth-sso/internals/domain/user"

type Repository interface {
    GetByID(id int) (*user.User, error)
    Create(u *user.User) error
}
package db

import (
	"auth-sso/internals/domain/user"
	userport "auth-sso/internals/ports/user"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) userport.Repository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(id int) (*user.User, error) {
    ctx := context.Background()

    var u user.User
    err := r.DB.QueryRow(
        ctx,
        "SELECT id, name, email FROM users WHERE id = $1",
        id,
    ).Scan(&u.ID, &u.Name, &u.Email)

    if err != nil {
        return nil, err
    }

    return &u, nil
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
    var createdUser user.User

    query := `
        INSERT INTO users (name, email)
        VALUES ($1, $2)
        RETURNING id, name, email
    `

    err := r.DB.QueryRow(
        ctx,
        query,
        u.Name,
        u.Email,
    ).Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email)

    if err != nil {
        return nil, err
    }

    return &createdUser, nil
}


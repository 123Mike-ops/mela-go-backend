package user

import (
	"auth-sso/internals/domain/user"
	userport "auth-sso/internals/ports/user"
	"context"
)

type Service struct {
    repo userport.Repository
}

func NewService(r userport.Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) GetUser(contet context.Context,id int) (*user.User, error) {
    return s.repo.GetByID(id)
}

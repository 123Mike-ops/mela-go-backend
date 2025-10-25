package user

import (
	"auth-sso/internals/domain/user"
	User "auth-sso/internals/domain/user"
	RegisterUser "auth-sso/internals/infrastructure/handler"
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
func (s *Service) CreateUser(contet context.Context,user *RegisterUser.RegisterUser) (*user.User, error) {
	 domainUser := &User.User{
        Name:  user.Name,
        Email: user.Email,
    }

    createdUser, err := s.repo.Create(contet, domainUser)
    if err != nil {
        return nil, err
    }

    return createdUser, nil
}

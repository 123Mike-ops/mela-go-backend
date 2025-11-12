package user

import (
	User "auth-sso/internals/domain/user"
	RegisterUser "auth-sso/internals/infrastructure/handler"
	userport "auth-sso/internals/ports/user"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
    repo userport.Repository
}

func NewService(r userport.Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) GetUser(contet context.Context,id int) (*User.User, error) {
    return s.repo.GetByID(id)
}
func (s *Service) CreateUser(ctx context.Context, user *RegisterUser.RegisterUser) (*User.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
  
	domainUser := &User.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdUser, err := s.repo.Create(ctx, domainUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}


func (s *Service) ValidateUserUniqueness(ctx context.Context, email, phoneNumber string) error {
    exists, err := s.repo.ExistsByEmailOrPhone(ctx, email, phoneNumber)
    if err != nil {
        return fmt.Errorf("failed to check uniqueness: %w", err)
    }
    if exists {
        return fmt.Errorf("email or phone number is already registered")
    }
    return nil
}
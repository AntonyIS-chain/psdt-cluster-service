package services

import (
	"context"
	"errors"

	"github.com/AntonyIS-chain/psdt-core-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-core-service/internal/core/ports"
)

type UserServiceImpl struct {
	userRepo ports.UserRepository
	// ldapClient ports.LDAPClient
}

func NewUserManagementService(userRepo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Authenticate(ctx context.Context, username, password string) (domain.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	// Placeholder password validation
	if user.Username != username || password == "" {
		return domain.User{}, errors.New("invalid username or password")
	}

	return user, nil
}

func (s *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}

func (s *UserServiceImpl) ListUsers(ctx context.Context) ([]domain.User, error) {
	return s.userRepo.ListUsers(ctx)
}

func (s *UserServiceImpl) InviteUser(ctx context.Context, email string) (string, error) {
	return s.userRepo.InviteUser(ctx, email)
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, signup domain.UserSignup) (domain.User, error) {
	// ldapUser, err := s.ldapClient.Authenticate(ctx, signup.Username, signup.Password)
	// if err != nil {
	// 	return domain.User{}, fmt.Errorf("LDAP authentication failed: %w", err)
	// }

	// ldapSignup := domain.User{
	// 	Username:  signup.Username,
	// 	Email:     ldapUser.Email,
	// 	FirstName: ldapUser.FirstName,
	// 	LastName:  ldapUser.LastName,
	// }

	user := domain.User{}
	// return s.userRepo.RegisterUser(ctx, ldapSignup)
	return s.userRepo.RegisterUser(ctx, user)
}

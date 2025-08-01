package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/google/uuid"
)

type UserManagementService struct {
	repo   ports.UserRepository
	logger ports.LoggingService
}

func NewUserManagementService(repo ports.UserRepository, logger ports.LoggingService) *UserManagementService {
	return &UserManagementService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserManagementService) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = user.CreatedAt

	if user.CreatedBy == "" {
		return "", fmt.Errorf("missing required field: CreatedBy")
	}

	_, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error("failed to create User", map[string]interface{}{
			"error": err.Error(),
		})
		return "", fmt.Errorf("create User: %w", err)
	}

	s.logger.Info("User created", map[string]interface{}{
		"id":         user.ID,
		"created_by": user.CreatedBy,
	})

	return user.ID, nil
}

func (s *UserManagementService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	dev, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error("get User failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return nil, fmt.Errorf("get User: %w", err)
	}
	return dev, nil
}

func (s *UserManagementService) ListUsers(ctx context.Context) ([]domain.User, error) {
	devs, err := s.repo.ListUsers(ctx)
	if err != nil {
		s.logger.Error("list Users failed", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("list Users: %w", err)
	}
	return devs, nil
}

func (s *UserManagementService) UpdateUser(ctx context.Context, id string, updates *domain.User) error {
	updates.UpdatedAt = time.Now().UTC()

	if updates.UpdatedBy == "" {
		return fmt.Errorf("missing required field: UpdatedBy")
	}

	err := s.repo.UpdateUser(ctx, id, updates)
	if err != nil {
		s.logger.Error("update User failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return fmt.Errorf("update User: %w", err)
	}

	s.logger.Info("User updated", map[string]interface{}{
		"id":         id,
		"updated_by": updates.UpdatedBy,
	})

	return nil
}

func (s *UserManagementService) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("delete User failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return fmt.Errorf("delete User: %w", err)
	}

	s.logger.Info("User deleted", map[string]interface{}{
		"id": id,
	})

	return nil
}

package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/google/uuid"
)

type SquadManagementService struct {
	repo   ports.SquadRepository
	logger ports.LoggingService
}

func NewSquadManagementService(repo ports.SquadRepository, logger ports.LoggingService) *SquadManagementService {
	return &SquadManagementService{
		repo:   repo,
		logger: logger,
	}
}

func (s *SquadManagementService) CreateSquad(ctx context.Context, squad *domain.Squad) (string, error) {
	s.logger.Info("Creating new squad", map[string]interface{}{"name": squad.Name})

	squad.ID = uuid.New().String()
	squad.CreatedAt = time.Now().UTC()
	squad.UpdatedAt = squad.CreatedAt

	id, err := s.repo.CreateSquad(ctx, squad)
	if err != nil {
		s.logger.Error("Failed to create squad", map[string]interface{}{"error": err.Error()})
		return "", fmt.Errorf("failed to create squad: %w", err)
	}

	s.logger.Info("Squad created successfully", map[string]interface{}{"squad_id": id})
	return id, nil
}

func (s *SquadManagementService) GetSquadByID(ctx context.Context, id string) (*domain.Squad, error) {
	s.logger.Info("Fetching squad by ID", map[string]interface{}{"squad_id": id})

	squad, err := s.repo.GetSquadByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to fetch squad", map[string]interface{}{"squad_id": id, "error": err.Error()})
		return nil, fmt.Errorf("failed to fetch squad: %w", err)
	}
	return squad, nil
}

func (s *SquadManagementService) ListSquads(ctx context.Context) ([]domain.Squad, error) {
	s.logger.Info("Listing all squads", nil)

	squads, err := s.repo.ListSquads(ctx)
	if err != nil {
		s.logger.Error("Failed to list squads", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to list squads: %w", err)
	}

	var result []domain.Squad
	for _, s := range squads {
		result = append(result, s)
	}

	s.logger.Info("Successfully listed squads", map[string]interface{}{"count": len(result)})
	return result, nil
}

func (s *SquadManagementService) UpdateSquad(ctx context.Context, id string, updates *domain.Squad) error {
	s.logger.Info("Updating squad", map[string]interface{}{"squad_id": id})

	updates.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateSquad(ctx, id, updates); err != nil {
		s.logger.Error("Failed to update squad", map[string]interface{}{"squad_id": id, "error": err.Error()})
		return fmt.Errorf("failed to update squad: %w", err)
	}

	s.logger.Info("Squad updated successfully", map[string]interface{}{"squad_id": id})
	return nil
}

func (s *SquadManagementService) DeleteSquad(ctx context.Context, id string) error {
	s.logger.Info("Deleting squad", map[string]interface{}{"squad_id": id})

	if err := s.repo.DeleteSquad(ctx, id); err != nil {
		s.logger.Error("Failed to delete squad", map[string]interface{}{"squad_id": id, "error": err.Error()})
		return fmt.Errorf("failed to delete squad: %w", err)
	}

	s.logger.Info("Squad deleted successfully", map[string]interface{}{"squad_id": id})
	return nil
}

func (s *SquadManagementService) AddUserToSquad(ctx context.Context, squadID, userID string) error {
	s.logger.Info("Adding user to squad", map[string]interface{}{"squad_id": squadID, "user_id": userID})

	err := s.repo.AddUserToSquad(ctx, squadID, userID)
	if err != nil {
		s.logger.Error("Failed to add user to squad", map[string]interface{}{
			"squad_id": squadID, "user_id": userID, "error": err.Error(),
		})
		return fmt.Errorf("failed to add user to squad: %w", err)
	}

	s.logger.Info("User added to squad successfully", map[string]interface{}{"squad_id": squadID, "user_id": userID})
	return nil
}

func (s *SquadManagementService) RemoveUserFromSquad(ctx context.Context, squadID, userID string) error {
	s.logger.Info("Removing user from squad", map[string]interface{}{"squad_id": squadID, "user_id": userID})

	err := s.repo.RemoveUserFromSquad(ctx, squadID, userID)
	if err != nil {
		s.logger.Error("Failed to remove user from squad", map[string]interface{}{
			"squad_id": squadID, "user_id": userID, "error": err.Error(),
		})
		return fmt.Errorf("failed to remove user from squad: %w", err)
	}

	s.logger.Info("User removed from squad successfully", map[string]interface{}{"squad_id": squadID, "user_id": userID})
	return nil
}

func (s *SquadManagementService) ListUsersInSquad(ctx context.Context, squadID string) ([]*domain.User, error) {
	s.logger.Info("Listing users in squad", map[string]interface{}{"squad_id": squadID})

	users, err := s.repo.ListUsersInSquad(ctx, squadID)
	if err != nil {
		s.logger.Error("Failed to list users in squad", map[string]interface{}{
			"squad_id": squadID, "error": err.Error(),
		})
		return nil, fmt.Errorf("failed to list users in squad: %w", err)
	}

	s.logger.Info("Successfully listed users in squad", map[string]interface{}{"squad_id": squadID, "user_count": len(users)})
	return users, nil
}

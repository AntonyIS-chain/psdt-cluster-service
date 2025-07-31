package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-core-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-core-service/internal/core/ports"
	"github.com/google/uuid"
)

type SquadManagementService struct {
	repo ports.SquadRepository
}

func NewSquadManagementService(repo ports.SquadRepository) *SquadManagementService {
	return &SquadManagementService{
		repo: repo,
	}
}

func (s *SquadManagementService) GetSquadByID(ctx context.Context, id string) (domain.Squad, error) {
	squad, err := s.repo.GetSquadByID(ctx, id)
	if err != nil {
		return domain.Squad{}, fmt.Errorf("failed to get squad by ID: %w", err)
	}
	return squad, nil
}

func (s *SquadManagementService) ListSquads(ctx context.Context, tribeID string) ([]domain.Squad, error) {
	squads, err := s.repo.ListSquads(ctx, tribeID)
	if err != nil {
		return nil, fmt.Errorf("failed to list squads: %w", err)
	}
	return squads, nil
}

func (s *SquadManagementService) CreateSquad(ctx context.Context, squad domain.Squad) (string, error) {
	squad.ID = uuid.New().String()
	squad.CreatedAt = time.Now().UTC()
	squad.UpdatedAt = squad.CreatedAt

	id, err := s.repo.CreateSquad(ctx, squad)
	if err != nil {
		return "", fmt.Errorf("failed to create squad: %w", err)
	}
	return id, nil
}

func (s *SquadManagementService) UpdateSquad(ctx context.Context, id string, updates domain.Squad) error {
	updates.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateSquad(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to update squad: %w", err)
	}
	return nil
}

func (s *SquadManagementService) DeleteSquad(ctx context.Context, id string) error {
	if err := s.repo.DeleteSquad(ctx, id); err != nil {
		return fmt.Errorf("failed to delete squad: %w", err)
	}
	return nil
}

func (s *SquadManagementService) AddUserToSquad(ctx context.Context, user domain.SquadUser) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().UTC()

	if err := s.repo.AddUserToSquad(ctx, user); err != nil {
		return fmt.Errorf("failed to add user to squad: %w", err)
	}
	return nil
}

func (s *SquadManagementService) RemoveUserFromSquad(ctx context.Context, userID, squadID string) error {
	if err := s.repo.RemoveUserFromSquad(ctx, userID, squadID); err != nil {
		return fmt.Errorf("failed to remove user from squad: %w", err)
	}
	return nil
}

func (s *SquadManagementService) ListUsersInSquad(ctx context.Context, squadID string) ([]domain.User, error) {
	users, err := s.repo.ListUsersInSquad(ctx, squadID)
	if err != nil {
		return nil, fmt.Errorf("failed to list users in squad: %w", err)
	}
	return users, nil
}

func (s *SquadManagementService) GetSquadsForUser(ctx context.Context, userID string) ([]domain.Squad, error) {
	squads, err := s.repo.GetSquadsForUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list squads for user: %w", err)
	}
	return squads, nil
}

package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-core-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-core-service/internal/core/ports"
	"github.com/google/uuid"
)

type TribeManagementService struct {
	repo ports.TribeRepository
}

func NewTribeManagementService(repo ports.TribeRepository) *TribeManagementService {
	return &TribeManagementService{
		repo: repo,
	}
}

func (s *TribeManagementService) GetTribeByID(ctx context.Context, id string) (domain.Tribe, error) {
	tribe, err := s.repo.GetTribeByID(ctx, id)
	if err != nil {
		return domain.Tribe{}, fmt.Errorf("failed to get tribe by ID: %w", err)
	}
	return tribe, nil
}

func (s *TribeManagementService) ListTribes(ctx context.Context, clusterID string) ([]domain.Tribe, error) {
	tribes, err := s.repo.ListTribes(ctx, clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to list tribes: %w", err)
	}
	return tribes, nil
}

func (s *TribeManagementService) CreateTribe(ctx context.Context, tribe domain.Tribe) (string, error) {
	tribe.ID = uuid.New().String()
	tribe.CreatedAt = time.Now().UTC()
	tribe.UpdatedAt = tribe.CreatedAt

	id, err := s.repo.CreateTribe(ctx, tribe)
	if err != nil {
		return "", fmt.Errorf("failed to create tribe: %w", err)
	}
	return id, nil
}

func (s *TribeManagementService) UpdateTribe(ctx context.Context, id string, updates domain.Tribe) error {
	updates.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateTribe(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to update tribe: %w", err)
	}
	return nil
}

func (s *TribeManagementService) DeleteTribe(ctx context.Context, id string) error {
	if err := s.repo.DeleteTribe(ctx, id); err != nil {
		return fmt.Errorf("failed to delete tribe: %w", err)
	}
	return nil
}

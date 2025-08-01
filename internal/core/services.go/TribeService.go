package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/google/uuid"
)

type TribeManagementService struct {
	repo   ports.TribeRepository
	logger ports.LoggingService
}

func NewTribeManagementService(repo ports.TribeRepository, logger ports.LoggingService) *TribeManagementService {
	return &TribeManagementService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TribeManagementService) CreateTribe(ctx context.Context, tribe *domain.Tribe) (string, error) {
	s.logger.Info("Creating new tribe", map[string]interface{}{"name": tribe.Name})

	tribe.ID = uuid.New().String()
	tribe.CreatedAt = time.Now().UTC()
	tribe.UpdatedAt = tribe.CreatedAt

	id, err := s.repo.CreateTribe(ctx, tribe)
	if err != nil {
		s.logger.Error("Failed to create tribe", map[string]interface{}{"error": err.Error()})
		return "", fmt.Errorf("failed to create tribe: %w", err)
	}

	s.logger.Info("Tribe created successfully", map[string]interface{}{"tribe_id": id})
	return id, nil
}

func (s *TribeManagementService) GetTribeByID(ctx context.Context, id string) (*domain.Tribe, error) {
	s.logger.Info("Fetching tribe by ID", map[string]interface{}{"tribe_id": id})

	tribe, err := s.repo.GetTribeByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to fetch tribe", map[string]interface{}{"tribe_id": id, "error": err.Error()})
		return nil, fmt.Errorf("failed to fetch tribe: %w", err)
	}
	return tribe, nil
}

func (s *TribeManagementService) ListTribes(ctx context.Context) ([]*domain.Tribe, error) {
	s.logger.Info("Listing all tribes", nil)

	tribes, err := s.repo.ListTribes(ctx)
	if err != nil {
		s.logger.Error("Failed to list tribes", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to list tribes: %w", err)
	}

	s.logger.Info("Successfully listed tribes", map[string]interface{}{"count": len(tribes)})
	return tribes, nil
}

func (s *TribeManagementService) UpdateTribe(ctx context.Context, id string, updates *domain.Tribe) error {
	s.logger.Info("Updating tribe", map[string]interface{}{"tribe_id": id})

	updates.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateTribe(ctx, id, updates); err != nil {
		s.logger.Error("Failed to update tribe", map[string]interface{}{"tribe_id": id, "error": err.Error()})
		return fmt.Errorf("failed to update tribe: %w", err)
	}

	s.logger.Info("Tribe updated successfully", map[string]interface{}{"tribe_id": id})
	return nil
}

func (s *TribeManagementService) DeleteTribe(ctx context.Context, id string) error {
	s.logger.Info("Deleting tribe", map[string]interface{}{"tribe_id": id})

	if err := s.repo.DeleteTribe(ctx, id); err != nil {
		s.logger.Error("Failed to delete tribe", map[string]interface{}{"tribe_id": id, "error": err.Error()})
		return fmt.Errorf("failed to delete tribe: %w", err)
	}

	s.logger.Info("Tribe deleted successfully", map[string]interface{}{"tribe_id": id})
	return nil
}

func (s *TribeManagementService) AddSquadToTribe(ctx context.Context, tribeID, squadID string) error {
	s.logger.Info("Adding squad to tribe", map[string]interface{}{"tribe_id": tribeID, "squad_id": squadID})

	err := s.repo.AddSquadToTribe(ctx, tribeID, squadID)
	if err != nil {
		s.logger.Error("Failed to add squad to tribe", map[string]interface{}{
			"tribe_id": tribeID, "squad_id": squadID, "error": err.Error(),
		})
		return fmt.Errorf("failed to add squad to tribe: %w", err)
	}

	s.logger.Info("Squad added to tribe successfully", map[string]interface{}{"tribe_id": tribeID, "squad_id": squadID})
	return nil
}

func (s *TribeManagementService) RemoveSquadFromTribe(ctx context.Context, tribeID, squadID string) error {
	s.logger.Info("Removing squad from tribe", map[string]interface{}{"tribe_id": tribeID, "squad_id": squadID})

	err := s.repo.RemoveSquadFromTribe(ctx, tribeID, squadID)
	if err != nil {
		s.logger.Error("Failed to remove squad from tribe", map[string]interface{}{
			"tribe_id": tribeID, "squad_id": squadID, "error": err.Error(),
		})
		return fmt.Errorf("failed to remove squad from tribe: %w", err)
	}

	s.logger.Info("Squad removed from tribe successfully", map[string]interface{}{"tribe_id": tribeID, "squad_id": squadID})
	return nil
}

func (s *TribeManagementService) ListSquadsInTribe(ctx context.Context, tribeID string) ([]*domain.Squad, error) {
	s.logger.Info("Listing squads in tribe", map[string]interface{}{"tribe_id": tribeID})

	squads, err := s.repo.ListSquadsInTribe(ctx, tribeID)
	if err != nil {
		s.logger.Error("Failed to list squads in tribe", map[string]interface{}{
			"tribe_id": tribeID, "error": err.Error(),
		})
		return nil, fmt.Errorf("failed to list squads in tribe: %w", err)
	}

	s.logger.Info("Successfully listed squads in tribe", map[string]interface{}{"tribe_id": tribeID, "count": len(squads)})
	return squads, nil
}

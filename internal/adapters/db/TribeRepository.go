package db

import (
	"context"
	"fmt"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"gorm.io/gorm"
)

type TribeRepositoryImpl struct {
	db *gorm.DB
}

func NewTribeRepository(db *gorm.DB) *TribeRepositoryImpl {
	return &TribeRepositoryImpl{db: db}
}

func (r *TribeRepositoryImpl) CreateTribe(ctx context.Context, tribe *domain.Tribe) (string, error) {
	if err := r.db.WithContext(ctx).Create(tribe).Error; err != nil {
		return "", fmt.Errorf("failed to create tribe: %w", err)
	}
	return tribe.ID, nil
}

func (r *TribeRepositoryImpl) UpdateTribe(ctx context.Context, id string, tribe *domain.Tribe) error {
	result := r.db.WithContext(ctx).Model(&domain.Tribe{}).Where("id = ?", id).Updates(tribe)
	if result.Error != nil {
		return fmt.Errorf("failed to update tribe: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *TribeRepositoryImpl) GetTribeByID(ctx context.Context, id string) (*domain.Tribe, error) {
	var tribe domain.Tribe
	if err := r.db.WithContext(ctx).First(&tribe, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch tribe by ID: %w", err)
	}
	return &tribe, nil
}

func (r *TribeRepositoryImpl) ListTribes(ctx context.Context) ([]*domain.Tribe, error) {
	var tribes []*domain.Tribe
	if err := r.db.WithContext(ctx).Find(&tribes).Error; err != nil {
		return nil, fmt.Errorf("failed to list tribes: %w", err)
	}
	return tribes, nil
}

func (r *TribeRepositoryImpl) DeleteTribe(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Tribe{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete tribe: %w", err)
	}
	return nil
}

func (r *TribeRepositoryImpl) AddSquadToTribe(ctx context.Context, tribeID, squadID string) error {
	result := r.db.WithContext(ctx).Model(&domain.Squad{}).Where("id = ?", squadID).Update("tribe_id", tribeID)
	if result.Error != nil {
		return fmt.Errorf("failed to add squad to tribe: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *TribeRepositoryImpl) RemoveSquadFromTribe(ctx context.Context, tribeID, squadID string) error {
	result := r.db.WithContext(ctx).Model(&domain.Squad{}).
		Where("id = ? AND tribe_id = ?", squadID, tribeID).
		Update("tribe_id", "")
	if result.Error != nil {
		return fmt.Errorf("failed to remove squad from tribe: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *TribeRepositoryImpl) ListSquadsInTribe(ctx context.Context, tribeID string) ([]*domain.Squad, error) {
	var squads []*domain.Squad
	if err := r.db.WithContext(ctx).Where("tribe_id = ?", tribeID).Find(&squads).Error; err != nil {
		return nil, fmt.Errorf("failed to list squads in tribe: %w", err)
	}
	return squads, nil
}

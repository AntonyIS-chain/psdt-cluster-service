package db

import (
	"context"
	"fmt"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"gorm.io/gorm"
)

type SquadRepositoryImpl struct {
	db *gorm.DB
}

func NewSquadRepository(db *gorm.DB) *SquadRepositoryImpl {
	return &SquadRepositoryImpl{db: db}
}

func (r *SquadRepositoryImpl) CreateSquad(ctx context.Context, squad *domain.Squad) (string, error) {
	if err := r.db.WithContext(ctx).Create(squad).Error; err != nil {
		return "", fmt.Errorf("failed to create squad: %w", err)
	}
	return squad.ID, nil
}

func (r *SquadRepositoryImpl) UpdateSquad(ctx context.Context, id string, squad *domain.Squad) error {
	result := r.db.WithContext(ctx).Model(&domain.Squad{}).Where("id = ?", id).Updates(squad)
	if result.Error != nil {
		return fmt.Errorf("failed to update squad: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *SquadRepositoryImpl) GetSquadByID(ctx context.Context, id string) (*domain.Squad, error) {
	var squad domain.Squad
	if err := r.db.WithContext(ctx).First(&squad, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get squad by ID: %w", err)
	}
	return &squad, nil
}

func (r *SquadRepositoryImpl) ListSquads(ctx context.Context) ([]domain.Squad, error) {
	var squads []domain.Squad
	if err := r.db.WithContext(ctx).Find(&squads).Error; err != nil {
		return nil, fmt.Errorf("failed to list squads: %w", err)
	}
	return squads, nil
}

func (r *SquadRepositoryImpl) DeleteSquad(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Squad{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete squad: %w", err)
	}
	return nil
}

func (r *SquadRepositoryImpl) AddUserToSquad(ctx context.Context, squadID, userID string) error {
	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Update("squad_id", squadID)
	if result.Error != nil {
		return fmt.Errorf("failed to add user to squad: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *SquadRepositoryImpl) RemoveUserFromSquad(ctx context.Context, squadID, userID string) error {
	result := r.db.WithContext(ctx).Model(&domain.User{}).
		Where("id = ? AND squad_id = ?", userID, squadID).
		Update("squad_id", "")
	if result.Error != nil {
		return fmt.Errorf("failed to remove user from squad: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *SquadRepositoryImpl) ListUsersInSquad(ctx context.Context, squadID string) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.WithContext(ctx).Where("squad_id = ?", squadID).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to list users in squad: %w", err)
	}
	return users, nil
}

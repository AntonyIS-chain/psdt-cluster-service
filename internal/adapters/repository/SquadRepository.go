package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/AntonyIS-chain/psdt-cluster-service/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type squadRepository struct {
	db     *gorm.DB
	logger ports.LoggingService
}

func NewSquadRepository(db *gorm.DB, logger ports.LoggingService) *squadRepository {
	return &squadRepository{db: db, logger: logger}
}


func (r *squadRepository) GetSquadByID(ctx context.Context, id string) (domain.Squad, error) {
	var squad domain.Squad
	op := "GetSquadByID"

	err := r.db.Preload("Users").First(&squad, "id = ?", id).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to get squad by ID: %v", err), map[string]string{"id": id})
		return domain.Squad{}, fmt.Errorf("squad not found: %w", err)
	}
	return squad, nil
}

func (r *squadRepository) ListSquads(ctx context.Context, tribeID string) ([]domain.Squad, error) {
	var squads []domain.Squad
	op := "ListSquads"

	err := r.db.Where("tribe_id = ?", tribeID).Find(&squads).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to list squads: %v", err), map[string]string{"tribe_id": tribeID})
		return nil, fmt.Errorf("failed to list squads: %w", err)
	}
	return squads, nil
}

func (r *squadRepository) CreateSquad(ctx context.Context, squad domain.Squad) (string, error) {
	op := "CreateSquad"

	squad.ID = uuid.NewString()
	squad.CreatedAt = time.Now()
	squad.UpdatedAt = time.Now()

	if err := r.db.Create(&squad).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to create squad: %v", err), nil)
		return "", fmt.Errorf("failed to create squad: %w", err)
	}
	r.logInfo(ctx, op, "Squad created successfully", map[string]string{"squad_id": squad.ID})
	return squad.ID, nil
}

func (r *squadRepository) UpdateSquad(ctx context.Context, id string, updates domain.Squad) error {
	op := "UpdateSquad"

	updates.UpdatedAt = time.Now()
	result := r.db.Model(&domain.Squad{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to update squad: %v", result.Error), map[string]string{"squad_id": id})
		return fmt.Errorf("failed to update squad: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		r.logError(ctx, op, "No squad found for update", map[string]string{"squad_id": id})
		return fmt.Errorf("no squad found with ID %s", id)
	}
	r.logInfo(ctx, op, "Squad updated successfully", map[string]string{"squad_id": id})
	return nil
}

func (r *squadRepository) DeleteSquad(ctx context.Context, id string) error {
	op := "DeleteSquad"

	if err := r.db.Delete(&domain.Squad{}, "id = ?", id).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to delete squad: %v", err), map[string]string{"squad_id": id})
		return fmt.Errorf("failed to delete squad: %w", err)
	}
	r.logInfo(ctx, op, "Squad deleted successfully", map[string]string{"squad_id": id})
	return nil
}

func (r *squadRepository) AddUserToSquad(ctx context.Context, user domain.SquadUser) error {
	op := "AddUserToSquad"

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := r.db.Create(&user).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to add user to squad: %v", err), map[string]string{
			"user_id":  user.UserID,
			"squad_id": user.SquadID,
		})
		return fmt.Errorf("failed to add user to squad: %w", err)
	}
	r.logInfo(ctx, op, "User added to squad", map[string]string{
		"user_id":  user.UserID,
		"squad_id": user.SquadID,
	})
	return nil
}

func (r *squadRepository) RemoveUserFromSquad(ctx context.Context, userID, squadID string) error {
	op := "RemoveUserFromSquad"

	if err := r.db.Where("user_id = ? AND squad_id = ?", userID, squadID).
		Delete(&domain.SquadUser{}).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to remove user from squad: %v", err), map[string]string{
			"user_id":  userID,
			"squad_id": squadID,
		})
		return fmt.Errorf("failed to remove user from squad: %w", err)
	}
	r.logInfo(ctx, op, "User removed from squad", map[string]string{
		"user_id":  userID,
		"squad_id": squadID,
	})
	return nil
}

func (r *squadRepository) ListUsersInSquad(ctx context.Context, squadID string) ([]domain.User, error) {
	op := "ListUsersInSquad"

	var users []domain.User
	err := r.db.
		Model(&domain.User{}).
		Joins("JOIN squad_users ON squad_users.user_id = users.id").
		Where("squad_users.squad_id = ?", squadID).
		Find(&users).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to list users in squad: %v", err), map[string]string{"squad_id": squadID})
		return nil, fmt.Errorf("failed to list users in squad: %w", err)
	}
	return users, nil
}

func (r *squadRepository) GetSquadsForUser(ctx context.Context, userID string) ([]domain.Squad, error) {
	op := "GetSquadsForUser"

	var squads []domain.Squad
	err := r.db.
		Model(&domain.Squad{}).
		Joins("JOIN squad_users ON squad_users.squad_id = squads.id").
		Where("squad_users.user_id = ?", userID).
		Find(&squads).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to get squads for user: %v", err), map[string]string{"user_id": userID})
		return nil, fmt.Errorf("failed to get squads for user: %w", err)
	}
	return squads, nil
}

func (r *squadRepository) logError(ctx context.Context, operation, message string, fields map[string]string) {
	r.logger.LogError(ctx, domain.LogData{
		Level:         "ERROR",
		Message:       message,
		ServiceName:   serviceName,
		Environment:   environment,
		Version:       version,
		RequestID:     pkg.GetRequestID(ctx),
		CorrelationID: pkg.GetCorrelationID(ctx),
		TraceID:       pkg.GetTraceID(ctx),
		Operation:     operation,
		Fields:        fields,
	})
}

func (r *squadRepository) logInfo(ctx context.Context, operation, message string, fields map[string]string) {
	r.logger.LogInfo(ctx, domain.LogData{
		Level:         "INFO",
		Message:       message,
		ServiceName:   serviceName,
		Environment:   environment,
		Version:       version,
		RequestID:     pkg.GetRequestID(ctx),
		CorrelationID: pkg.GetCorrelationID(ctx),
		TraceID:       pkg.GetTraceID(ctx),
		Operation:     operation,
		Fields:        fields,
	})
}

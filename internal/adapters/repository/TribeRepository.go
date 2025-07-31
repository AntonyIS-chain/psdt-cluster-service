package repository

import (
	"context"
	"fmt"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/AntonyIS-chain/psdt-cluster-service/pkg"
	"gorm.io/gorm"
)

type TribeRepositoryImpl struct {
	db     *gorm.DB
	logger ports.LoggingService
}

func NewTribeRepository(db *gorm.DB, logger ports.LoggingService) *TribeRepositoryImpl {
	return &TribeRepositoryImpl{db: db, logger: logger}
}

func (r *TribeRepositoryImpl) GetTribeByID(ctx context.Context, id string) (domain.Tribe, error) {
	var tribe domain.Tribe
	op := "GetTribeByID"

	if err := r.db.First(&tribe, "id = ?", id).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to fetch tribe by ID: %v", err), map[string]string{
			"tribe_id": id,
		})
		return domain.Tribe{}, err
	}

	r.logInfo(ctx, op, "Fetched tribe by ID", map[string]string{"tribe_id": id})
	return tribe, nil
}

func (r *TribeRepositoryImpl) ListTribes(ctx context.Context, clusterID string) ([]domain.Tribe, error) {
	var tribes []domain.Tribe
	op := "ListTribes"

	if err := r.db.Where("cluster_id = ?", clusterID).Find(&tribes).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to list tribes by cluster ID: %v", err), map[string]string{
			"cluster_id": clusterID,
		})
		return nil, err
	}

	r.logInfo(ctx, op, "Listed tribes by cluster ID", map[string]string{
		"cluster_id": clusterID,
		"count":      fmt.Sprintf("%d", len(tribes)),
	})
	return tribes, nil
}

func (r *TribeRepositoryImpl) CreateTribe(ctx context.Context, tribe domain.Tribe) (string, error) {
	op := "CreateTribe"

	if err := r.db.Create(&tribe).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to create tribe: %v", err), map[string]string{
			"name": tribe.Name,
		})
		return "", err
	}

	r.logInfo(ctx, op, "Created tribe successfully", map[string]string{
		"tribe_id": tribe.ID,
		"name":     tribe.Name,
	})
	return tribe.ID, nil
}

func (r *TribeRepositoryImpl) UpdateTribe(ctx context.Context, id string, updates domain.Tribe) error {
	op := "UpdateTribe"

	if err := r.db.Model(&domain.Tribe{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to update tribe: %v", err), map[string]string{
			"tribe_id": id,
		})
		return err
	}

	r.logInfo(ctx, op, "Updated tribe successfully", map[string]string{"tribe_id": id})
	return nil
}

func (r *TribeRepositoryImpl) DeleteTribe(ctx context.Context, id string) error {
	op := "DeleteTribe"

	if err := r.db.Delete(&domain.Tribe{}, "id = ?", id).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to delete tribe: %v", err), map[string]string{
			"tribe_id": id,
		})
		return err
	}

	r.logInfo(ctx, op, "Deleted tribe successfully", map[string]string{"tribe_id": id})
	return nil
}

// --- Structured logging helpers ---

func (r *TribeRepositoryImpl) logError(ctx context.Context, operation, message string, fields map[string]string) {
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

func (r *TribeRepositoryImpl) logInfo(ctx context.Context, operation, message string, fields map[string]string) {
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

package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/AntonyIS-chain/psdt-cluster-service/pkg"
	"gorm.io/gorm"
)

type ClusterRepositoryImpl struct {
	db     *gorm.DB
	logger ports.LoggingService
}

func NewClusterRepository(db *gorm.DB, logger ports.LoggingService) *ClusterRepositoryImpl {
	return &ClusterRepositoryImpl{db: db, logger: logger}
}

var (
	serviceName = "psdt-cluster-service"
	environment = os.Getenv("ENVIRONMENT") // e.g., "dev", "prod"
	version     = os.Getenv("VERSION")     // e.g., "v1.0.0"
)

func (r *ClusterRepositoryImpl) GetClusterByID(ctx context.Context, id string) (domain.Cluster, error) {
	var cluster domain.Cluster
	op := "GetClusterByID"

	if err := r.db.First(&cluster, "id = ?", id).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to get cluster by ID %s: %v", id, err), map[string]string{
			"cluster_id": id,
		})
		return domain.Cluster{}, err
	}
	return cluster, nil
}

func (r *ClusterRepositoryImpl) ListClusters(ctx context.Context) ([]domain.Cluster, error) {
	var clusters []domain.Cluster
	op := "ListClusters"

	if err := r.db.Find(&clusters).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to list clusters: %v", err), nil)
		return nil, err
	}
	return clusters, nil
}

func (r *ClusterRepositoryImpl) CreateCluster(ctx context.Context, cluster domain.Cluster) (string, error) {
	op := "CreateCluster"
	if err := r.db.Create(&cluster).Error; err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to create cluster %+v: %v", cluster, err), nil)
		return "", err
	}

	r.logInfo(ctx, op, fmt.Sprintf("Cluster created successfully with ID %s", cluster.ID), map[string]string{
		"cluster_id": cluster.ID,
	})
	return cluster.ID, nil
}

func (r *ClusterRepositoryImpl) UpdateCluster(ctx context.Context, id string, updates domain.Cluster) error {
	op := "UpdateCluster"
	err := r.db.Model(&domain.Cluster{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to update cluster %s: %v", id, err), map[string]string{
			"cluster_id": id,
		})
	}
	return err
}

func (r *ClusterRepositoryImpl) DeleteCluster(ctx context.Context, id string) error {
	op := "DeleteCluster"
	err := r.db.Delete(&domain.Cluster{}, "id = ?", id).Error
	if err != nil {
		r.logError(ctx, op, fmt.Sprintf("Failed to delete cluster with ID %s: %v", id, err), map[string]string{
			"cluster_id": id,
		})
	}
	return err
}


func (r *ClusterRepositoryImpl) logError(ctx context.Context, operation, message string, fields map[string]string) {
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

func (r *ClusterRepositoryImpl) logInfo(ctx context.Context, operation, message string, fields map[string]string) {
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

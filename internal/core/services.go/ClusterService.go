package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/google/uuid"
)

type ClusterManagementService struct {
	repo   ports.ClusterRepository
	logger ports.LoggingService
}

func NewClusterManagementService(repo ports.ClusterRepository, logger ports.LoggingService) *ClusterManagementService {
	return &ClusterManagementService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ClusterManagementService) CreateCluster(ctx context.Context, cluster *domain.Cluster) (string, error) {
	cluster.ID = uuid.New().String()
	cluster.CreatedAt = time.Now().UTC()
	cluster.UpdatedAt = cluster.CreatedAt

	if cluster.CreatedBy == "" {
		return "", fmt.Errorf("missing required field: CreatedBy")
	}

	id, err := s.repo.CreateCluster(ctx, cluster)
	if err != nil {
		s.logger.Error("failed to create cluster", map[string]interface{}{
			"error": err.Error(),
		})
		return "", fmt.Errorf("create cluster: %w", err)
	}

	s.logger.Info("cluster created", map[string]interface{}{
		"id":         id,
		"created_by": cluster.CreatedBy,
	})

	return id, nil
}

func (s *ClusterManagementService) GetClusterByID(ctx context.Context, id string) (*domain.Cluster, error) {
	cluster, err := s.repo.GetClusterByID(ctx, id)
	if err != nil {
		s.logger.Error("get cluster failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return nil, fmt.Errorf("get cluster: %w", err)
	}
	return cluster, nil
}

func (s *ClusterManagementService) ListClusters(ctx context.Context) ([]*domain.Cluster, error) {
	clusters, err := s.repo.ListClusters(ctx)
	if err != nil {
		s.logger.Error("list clusters failed", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("list clusters: %w", err)
	}
	return clusters, nil
}

func (s *ClusterManagementService) UpdateCluster(ctx context.Context, id string, updates *domain.Cluster) error {
	updates.UpdatedAt = time.Now().UTC()

	if updates.UpdatedBy == "" {
		return fmt.Errorf("missing required field: UpdatedBy")
	}

	err := s.repo.UpdateCluster(ctx, id, updates)
	if err != nil {
		s.logger.Error("update cluster failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return fmt.Errorf("update cluster: %w", err)
	}

	s.logger.Info("cluster updated", map[string]interface{}{
		"id":         id,
		"updated_by": updates.UpdatedBy,
	})

	return nil
}

func (s *ClusterManagementService) DeleteCluster(ctx context.Context, id string) error {
	err := s.repo.DeleteCluster(ctx, id)
	if err != nil {
		s.logger.Error("delete cluster failed", map[string]interface{}{
			"error": err.Error(),
			"id":    id,
		})
		return fmt.Errorf("delete cluster: %w", err)
	}

	s.logger.Info("cluster deleted", map[string]interface{}{
		"id": id,
	})

	return nil
}

func (s *ClusterManagementService) AddTribeToCluster(ctx context.Context, clusterID, tribeID string) error {
	if clusterID == "" || tribeID == "" {
		return fmt.Errorf("missing required clusterID or tribeID")
	}

	err := s.repo.AddTribeToCluster(ctx, clusterID, tribeID)
	if err != nil {
		s.logger.Error("failed to add tribe to cluster", map[string]interface{}{
			"error":     err.Error(),
			"clusterID": clusterID,
			"tribeID":   tribeID,
		})
		return fmt.Errorf("add tribe to cluster: %w", err)
	}

	s.logger.Info("tribe added to cluster", map[string]interface{}{
		"clusterID": clusterID,
		"tribeID":   tribeID,
	})

	return nil
}

func (s *ClusterManagementService) RemoveTribeFromCluster(ctx context.Context, clusterID, tribeID string) error {
	if clusterID == "" || tribeID == "" {
		return fmt.Errorf("missing required clusterID or tribeID")
	}

	err := s.repo.RemoveTribeFromCluster(ctx, clusterID, tribeID)
	if err != nil {
		s.logger.Error("failed to remove tribe from cluster", map[string]interface{}{
			"error":     err.Error(),
			"clusterID": clusterID,
			"tribeID":   tribeID,
		})
		return fmt.Errorf("remove tribe from cluster: %w", err)
	}

	s.logger.Info("tribe removed from cluster", map[string]interface{}{
		"clusterID": clusterID,
		"tribeID":   tribeID,
	})

	return nil
}

func (s *ClusterManagementService) ListTribesInCluster(ctx context.Context, clusterID string) ([]*domain.Tribe, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("missing required clusterID")
	}

	tribes, err := s.repo.ListTribesInCluster(ctx, clusterID)
	if err != nil {
		s.logger.Error("failed to list tribes in cluster", map[string]interface{}{
			"error":     err.Error(),
			"clusterID": clusterID,
		})
		return nil, fmt.Errorf("list tribes in cluster: %w", err)
	}

	return tribes, nil
}

package services

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-core-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-core-service/internal/core/ports"
	"github.com/google/uuid"
)

type ClusterManagementService struct {
	repo ports.ClusterRepository
}

func NewClusterManagementService(repo ports.ClusterRepository) *ClusterManagementService {
	return &ClusterManagementService{
		repo: repo,
	}
}

func (s *ClusterManagementService) CreateCluster(ctx context.Context, cluster domain.Cluster) (string, error) {
	cluster.ID = uuid.New().String()
	cluster.CreatedAt = time.Now().UTC()
	cluster.UpdatedAt = cluster.CreatedAt

	id, err := s.repo.CreateCluster(ctx, cluster)
	if err != nil {
		return "", fmt.Errorf("failed to create cluster: %w", err)
	}
	return id, nil
}

func (s *ClusterManagementService) GetClusterByID(ctx context.Context, id string) (domain.Cluster, error) {
	cluster, err := s.repo.GetClusterByID(ctx, id)
	if err != nil {
		return domain.Cluster{}, fmt.Errorf("failed to fetch cluster: %w", err)
	}
	return cluster, nil
}

func (s *ClusterManagementService) ListClusters(ctx context.Context) ([]domain.Cluster, error) {
	clusters, err := s.repo.ListClusters(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list clusters: %w", err)
	}
	return clusters, nil
}

func (s *ClusterManagementService) UpdateCluster(ctx context.Context, id string, updates domain.Cluster) error {
	updates.UpdatedAt = time.Now().UTC()

	if err := s.repo.UpdateCluster(ctx, id, updates); err != nil {
		return fmt.Errorf("failed to update cluster: %w", err)
	}
	return nil
}

func (s *ClusterManagementService) DeleteCluster(ctx context.Context, id string) error {
	if err := s.repo.DeleteCluster(ctx, id); err != nil {
		return fmt.Errorf("failed to delete cluster: %w", err)
	}
	return nil
}

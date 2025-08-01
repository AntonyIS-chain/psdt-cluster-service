package db

import (
	"context"
	"fmt"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"gorm.io/gorm"
)

type ClusterRepositoryImpl struct {
	db *gorm.DB
}

func NewClusterRepository(db *gorm.DB) *ClusterRepositoryImpl {
	return &ClusterRepositoryImpl{db: db}
}

func (r *ClusterRepositoryImpl) CreateCluster(ctx context.Context, cluster *domain.Cluster) (string, error) {
	if err := r.db.WithContext(ctx).Create(cluster).Error; err != nil {
		return "", fmt.Errorf("failed to create cluster: %w", err)
	}
	return cluster.ID, nil
}

func (r *ClusterRepositoryImpl) UpdateCluster(ctx context.Context, id string, cluster *domain.Cluster) error {
	result := r.db.WithContext(ctx).Model(&domain.Cluster{}).Where("id = ?", id).Updates(cluster)
	if result.Error != nil {
		return fmt.Errorf("failed to update cluster: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ClusterRepositoryImpl) GetClusterByID(ctx context.Context, id string) (*domain.Cluster, error) {
	var cluster domain.Cluster
	if err := r.db.WithContext(ctx).First(&cluster, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch cluster by ID: %w", err)
	}
	return &cluster, nil
}

func (r *ClusterRepositoryImpl) ListClusters(ctx context.Context) ([]*domain.Cluster, error) {
	var clusters []*domain.Cluster
	if err := r.db.WithContext(ctx).Find(&clusters).Error; err != nil {
		return nil, fmt.Errorf("failed to list clusters: %w", err)
	}
	return clusters, nil
}

func (r *ClusterRepositoryImpl) DeleteCluster(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Cluster{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete cluster: %w", err)
	}
	return nil
}

func (r *ClusterRepositoryImpl) AddTribeToCluster(ctx context.Context, clusterID, tribeID string) error {
	// Just update the Tribe's ClusterID
	result := r.db.WithContext(ctx).Model(&domain.Tribe{}).Where("id = ?", tribeID).Update("cluster_id", clusterID)
	if result.Error != nil {
		return fmt.Errorf("failed to add tribe to cluster: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ClusterRepositoryImpl) RemoveTribeFromCluster(ctx context.Context, clusterID, tribeID string) error {
	// Set Tribe's ClusterID to empty where both clusterID and tribeID match
	result := r.db.WithContext(ctx).Model(&domain.Tribe{}).
		Where("id = ? AND cluster_id = ?", tribeID, clusterID).
		Update("cluster_id", "")
	if result.Error != nil {
		return fmt.Errorf("failed to remove tribe from cluster: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ClusterRepositoryImpl) ListTribesInCluster(ctx context.Context, clusterID string) ([]*domain.Tribe, error) {
	var tribes []*domain.Tribe
	if err := r.db.WithContext(ctx).Where("cluster_id = ?", clusterID).Find(&tribes).Error; err != nil {
		return nil, fmt.Errorf("failed to list tribes in cluster: %w", err)
	}
	return tribes, nil
}

package ports

import (
	"context"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
)

// DeveloperRepository defines DB operations for Developer
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	UpdateUser(ctx context.Context, id string, user *domain.User) error
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}

// SquadRepository defines DB operations for Squad
type SquadRepository interface {
	CreateSquad(ctx context.Context, squad *domain.Squad) (string, error)
	UpdateSquad(ctx context.Context, id string, squad *domain.Squad) error
	GetSquadByID(ctx context.Context, id string) (*domain.Squad, error)
	ListSquads(ctx context.Context) ([]domain.Squad, error)
	DeleteSquad(ctx context.Context, id string) error
	AddUserToSquad(ctx context.Context, squadID, userID string) error
	RemoveUserFromSquad(ctx context.Context, squadID, userID string) error
	ListUsersInSquad(ctx context.Context, squadID string) ([]*domain.User, error)
}

// TribeRepository defines DB operations for Tribe
type TribeRepository interface {
	CreateTribe(ctx context.Context, tribe *domain.Tribe) (string, error)
	UpdateTribe(ctx context.Context, id string, tribe *domain.Tribe) error
	GetTribeByID(ctx context.Context, id string) (*domain.Tribe, error)
	ListTribes(ctx context.Context) ([]*domain.Tribe, error)
	DeleteTribe(ctx context.Context, id string) error
	AddSquadToTribe(ctx context.Context, tribeID, squadID string) error
	RemoveSquadFromTribe(ctx context.Context, tribeID, squadID string) error
	ListSquadsInTribe(ctx context.Context, tribeID string) ([]*domain.Squad, error)
}

// ClusterRepository defines DB operations for Cluster
type ClusterRepository interface {
	CreateCluster(ctx context.Context, cluster *domain.Cluster) (string, error)
	UpdateCluster(ctx context.Context, id string, cluster *domain.Cluster) error
	GetClusterByID(ctx context.Context, id string) (*domain.Cluster, error)
	ListClusters(ctx context.Context) ([]*domain.Cluster, error)
	DeleteCluster(ctx context.Context, id string) error
	AddTribeToCluster(ctx context.Context, clusterID, tribeID string) error
	RemoveTribeFromCluster(ctx context.Context, clusterID, tribeID string) error
	ListTribesInCluster(ctx context.Context, clusterID string) ([]*domain.Tribe, error)
}

type LoggingService interface {
	Info(message string, fields map[string]interface{})
	Error(message string, fields map[string]interface{})
	Debug(message string, fields map[string]interface{})
	Warn(message string, fields map[string]interface{})
}

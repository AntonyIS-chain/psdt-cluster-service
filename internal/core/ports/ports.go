package ports

import (
	"context"

	"github.com/AntonyIS-chain/psdt-core-service/internal/core/domain"
)

type ClusterService interface {
	GetClusterByID(ctx context.Context, id string) (domain.Cluster, error)
	ListClusters(ctx context.Context) ([]domain.Cluster, error)
	CreateCluster(ctx context.Context, cluster domain.Cluster) (string, error)
	UpdateCluster(ctx context.Context, id string, updates domain.Cluster) error
	DeleteCluster(ctx context.Context, id string) error
}

type ClusterRepository interface {
	GetClusterByID(ctx context.Context, id string) (domain.Cluster, error)
	ListClusters(ctx context.Context) ([]domain.Cluster, error)
	CreateCluster(ctx context.Context, cluster domain.Cluster) (string, error)
	UpdateCluster(ctx context.Context, id string, updates domain.Cluster) error
	DeleteCluster(ctx context.Context, id string) error
}

type TribeService interface {
	GetTribeByID(ctx context.Context, id string) (domain.Tribe, error)
	ListTribes(ctx context.Context, clusterID string) ([]domain.Tribe, error)
	CreateTribe(ctx context.Context, tribe domain.Tribe) (string, error)
	UpdateTribe(ctx context.Context, id string, updates domain.Tribe) error
	DeleteTribe(ctx context.Context, id string) error
}

type TribeRepository interface {
	GetTribeByID(ctx context.Context, id string) (domain.Tribe, error)
	ListTribes(ctx context.Context, clusterID string) ([]domain.Tribe, error)
	CreateTribe(ctx context.Context, tribe domain.Tribe) (string, error)
	UpdateTribe(ctx context.Context, id string, updates domain.Tribe) error
	DeleteTribe(ctx context.Context, id string) error
}

type SquadService interface {
	GetSquadByID(ctx context.Context, id string) (domain.Squad, error)
	ListSquads(ctx context.Context, tribeID string) ([]domain.Squad, error)
	CreateSquad(ctx context.Context, squad domain.Squad) (string, error)
	UpdateSquad(ctx context.Context, id string, updates domain.Squad) error
	DeleteSquad(ctx context.Context, id string) error
	AddUserToSquad(ctx context.Context, user domain.SquadUser) error
	RemoveUserFromSquad(ctx context.Context, userID, squadID string) error
	ListUsersInSquad(ctx context.Context, squadID string) ([]domain.User, error)
	GetSquadsForUser(ctx context.Context, userID string) ([]domain.Squad, error)
}

type SquadRepository interface {
	GetSquadByID(ctx context.Context, id string) (domain.Squad, error)
	ListSquads(ctx context.Context, tribeID string) ([]domain.Squad, error)
	CreateSquad(ctx context.Context, squad domain.Squad) (string, error)
	UpdateSquad(ctx context.Context, id string, updates domain.Squad) error
	DeleteSquad(ctx context.Context, id string) error
	AddUserToSquad(ctx context.Context, user domain.SquadUser) error
	RemoveUserFromSquad(ctx context.Context, userID, squadID string) error
	ListUsersInSquad(ctx context.Context, squadID string) ([]domain.User, error)
	GetSquadsForUser(ctx context.Context, userID string) ([]domain.Squad, error)
}

type UserService interface {
	Authenticate(ctx context.Context, username, password string) (domain.User, error)
	GetUserByUsername(ctx context.Context, email string) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	InviteUser(ctx context.Context, email string) (string, error)
	RegisterUser(ctx context.Context, signup domain.UserSignup) (domain.User, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, email string) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	InviteUser(ctx context.Context, email string) (string, error)
	RegisterUser(ctx context.Context, signup domain.User) (domain.User, error)
}

type LoggingService interface {
	LogInfo(ctx context.Context, data domain.LogData)
	LogError(ctx context.Context, data domain.LogData)
}

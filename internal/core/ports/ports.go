package ports

import "github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"

// DeveloperRepository defines DB operations for Developer
type DeveloperRepository interface {
	CreateDeveloper(dev *domain.Developer) error
	UpdateDeveloper(dev *domain.Developer) error
	GetDeveloperByID(id string) (*domain.Developer, error)
	ListDevelopers() ([]*domain.Developer, error)
	DeleteDeveloper(id string) error
}

// SquadRepository defines DB operations for Squad
type SquadRepository interface {
	CreateSquad(squad *domain.Squad) error
	UpdateSquad(squad *domain.Squad) error
	GetSquadByID(id string) (*domain.Squad, error)
	ListSquads() ([]*domain.Squad, error)
	DeleteSquad(id string) error
}

// TribeRepository defines DB operations for Tribe
type TribeRepository interface {
	CreateTribe(tribe *domain.Tribe) error
	UpdateTribe(tribe *domain.Tribe) error
	GetTribeByID(id string) (*domain.Tribe, error)
	ListTribes() ([]*domain.Tribe, error)
	DeleteTribe(id string) error
}

// ClusterRepository defines DB operations for Cluster
type ClusterRepository interface {
	CreateCluster(cluster *domain.Cluster) error
	UpdateCluster(cluster *domain.Cluster) error
	GetClusterByID(id string) (*domain.Cluster, error)
	ListClusters() ([]*domain.Cluster, error)
	DeleteCluster(id string) error
}

type LoggingService interface {
    Info(message string, fields map[string]interface{})
    Error(message string, fields map[string]interface{})
    Debug(message string, fields map[string]interface{})
    Warn(message string, fields map[string]interface{})
}

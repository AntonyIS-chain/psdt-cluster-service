package domain

import "time"

// Cluster represents a group of tribes
type Cluster struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Tribe belongs to a Cluster
type Tribe struct {
	ID          string    `json:"id" db:"id"`
	ClusterID   string    `json:"cluster_id" db:"cluster_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Squad belongs to a Tribe (and indirectly a Cluster)
type Squad struct {
	ID          string    `json:"id" db:"id"`
	TribeID     string    `json:"tribe_id" db:"tribe_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// User represents a developer or team member
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Username  string
	Role      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSignup struct {
	Username    string
	Password    string
	InviteToken string
}

// SquadUser is a many-to-many relationship between Users and Squads
type SquadUser struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	SquadID   string    `json:"squad_id" db:"squad_id"`
	Role      string    `json:"role" db:"role"` // Engineer, PM, TL, etc.
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type LogData struct {
	Level         string
	Message       string
	ServiceName   string
	Environment   string
	Version       string
	RequestID     string
	CorrelationID string
	TraceID       string
	Operation     string
	Fields        map[string]string
}

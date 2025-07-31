package domain

import "time"

type Cluster struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tribe struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ClusterID   string    `json:"cluster_id"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Squad struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	TribeID     string    `json:"tribe_id"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Developer struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	SquadID   string    `json:"squad_id"`
	IsActive  bool      `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



type LogData struct {
    LogID         string                 `json:"log_id"`          // Unique ID for log entry
    Timestamp     time.Time              `json:"timestamp"`       // Time the log was created
    Level         string                 `json:"level"`           // INFO, ERROR, DEBUG, etc.
    Message       string                 `json:"message"`         // Human-readable message
    Fields        map[string]interface{} `json:"fields"`          // Additional structured metadata

    // Contextual Metadata
    ServiceName   string                 `json:"service_name"`    // Which service generated it
    ServiceVersion string                `json:"service_version"` // Deployed version or git hash
    Environment   string                 `json:"environment"`     // e.g., production, staging
    Hostname      string                 `json:"hostname"`        // Host or pod name
    InstanceID    string                 `json:"instance_id"`     // Instance or container ID
    TraceID       string                 `json:"trace_id"`        // For tracing logs
    SpanID        string                 `json:"span_id"`         // For tracing spans
    CorrelationID string                 `json:"correlation_id"`  // To tie logs across services
    UserID        string                 `json:"user_id,omitempty"` // If action is user-specific
    Tags          []string               `json:"tags,omitempty"`    // Optional tags for filtering
}
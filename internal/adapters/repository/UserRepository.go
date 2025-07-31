package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/AntonyIS-chain/psdt-cluster-service/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type userRepository struct {
	db     *gorm.DB
	logger ports.LoggingService
}

func NewUserRepository(db *gorm.DB, logger ports.LoggingService) *userRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	op := "GetUserByUsername"

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		r.logError(ctx, op, "Failed to fetch user by email", map[string]string{
			"email": email,
			"error": err.Error(),
		})
		return domain.User{}, fmt.Errorf("user not found: %w", err)
	}

	r.logInfo(ctx, op, "Fetched user by email", map[string]string{
		"email": user.Email,
		"id":    user.ID,
	})
	return user, nil
}

func (r *userRepository) ListUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	op := "ListUsers"

	if err := r.db.Find(&users).Error; err != nil {
		r.logError(ctx, op, "Failed to list users", map[string]string{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	r.logInfo(ctx, op, "Listed users", map[string]string{
		"count": fmt.Sprintf("%d", len(users)),
	})
	return users, nil
}

func (r *userRepository) InviteUser(ctx context.Context, email string) (string, error) {
	op := "InviteUser"

	var existing domain.User
	if err := r.db.First(&existing, "email = ?", email).Error; err == nil {
		r.logError(ctx, op, "User already exists", map[string]string{
			"email": email,
		})
		return "", errors.New("user already exists")
	}

	inviteToken := uuid.NewString()

	user := domain.User{
		ID:        uuid.NewString(),
		Email:     email,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.db.Create(&user).Error; err != nil {
		r.logError(ctx, op, "Failed to create user (invite)", map[string]string{
			"email": email,
			"error": err.Error(),
		})
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	r.logInfo(ctx, op, "Invited user", map[string]string{
		"user_id":     user.ID,
		"email":       email,
		"inviteToken": inviteToken,
	})
	return inviteToken, nil
}

func (r *userRepository) RegisterUser(ctx context.Context, signup domain.User) (domain.User, error) {
	op := "RegisterUser"

	var existing domain.User
	if err := r.db.First(&existing, "email = ?", signup.Email).Error; err == nil {
		r.logError(ctx, op, "User already exists", map[string]string{
			"email": signup.Email,
		})
		return domain.User{}, errors.New("user already exists")
	}

	signup.ID = uuid.NewString()
	signup.CreatedAt = time.Now()
	signup.UpdatedAt = time.Now()
	signup.Active = true

	if err := r.db.Create(&signup).Error; err != nil {
		r.logError(ctx, op, "Failed to register user", map[string]string{
			"email": signup.Email,
			"error": err.Error(),
		})
		return domain.User{}, fmt.Errorf("failed to register user: %w", err)
	}

	r.logInfo(ctx, op, "Registered user", map[string]string{
		"user_id": signup.ID,
		"email":   signup.Email,
	})
	return signup, nil
}

// --- Structured logging helpers ---

func (r *userRepository) logError(ctx context.Context, operation, message string, fields map[string]string) {
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

func (r *userRepository) logInfo(ctx context.Context, operation, message string, fields map[string]string) {
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

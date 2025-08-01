package db

import (
	"context"
	"fmt"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return user.ID, nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, id string, user *domain.User) error {
	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID: %w", err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) ListUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

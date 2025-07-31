package repository

import (
	"fmt"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-core-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBClient struct {
	DB *gorm.DB
}

// NewDBClient initializes and returns a PostgreSQL GORM client.
func NewDBClient(cfg *config.Config) (*DBClient, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		cfg.POSTGRES_HOST, cfg.POSTGRES_PORT,
		cfg.POSTGRES_USER, cfg.POSTGRES_DB, cfg.POSTGRES_PASSWORD,
	)

	// Use silent logger to avoid logging SQL statements and caller paths
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, fmt.Errorf("could not open DB: %w", err)
	}

	// Setup connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql.DB from gorm DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Auto-migrate all required domain models
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Cluster{},
		&domain.Tribe{},
		&domain.Squad{},
		&domain.SquadUser{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed auto migration: %w", err)
	}

	return &DBClient{DB: db}, nil
}

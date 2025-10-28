package database

import (
	"fmt"
	"log"
	"time"

	"github.com/josevitorrodriguess/whisper/server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDBConfig() *DatabaseConfig {

	return &DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "whisper_user",
		Password: "whisper_password",
		DBName:   "whisper_db",
	}
}

func ConnectDatabase() *gorm.DB {
	config := loadDBConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	if err = runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to obtain database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected Postgres Successfully!")
	return db
}

func TestConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func runMigrations(db *gorm.DB) error {
	log.Println("Running Migrations...")
	if err := db.AutoMigrate(
		models.User{},
	); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

package database

import (
	"log"
	"os"

	"github.com/ctf-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Get DATABASE_URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// AutoMigrate will create tables if they don’t exist
	err = db.AutoMigrate(
		&models.Team{},
		&models.Challenge{},
		&models.Submission{},
		&models.Score{},
	)
	if err != nil {
		log.Fatalf("❌ AutoMigration failed: %v", err)
	}

	log.Println("✅ Database connected and migrated successfully")
	return db
}

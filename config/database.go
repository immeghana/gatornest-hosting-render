package config

import (
	"backend/internal/entity"
	"backend/internal/migration"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=gatornest port=5432 sslmode=disable"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Run migrations
	if err := migration.RunMigrations(); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
	}

	// Auto-migrate with new fields
	err = DB.AutoMigrate(&entity.Student{}, &entity.Room{}, &entity.Admin{}, &entity.MaintenanceRequest{}, &entity.Payment{})

	if err != nil {
		panic("Failed to migrate database!")
	}

	fmt.Println("Database connected successfully!")
}

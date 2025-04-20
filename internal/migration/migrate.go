package migration

import (
	"backend/config"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RunMigrations() error {
	db := config.DB

	// Read the SQL file
	sqlFile := filepath.Join("internal", "migration", "0001_alter_student_table.sql")
	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	// Split the SQL file into individual statements
	sqlStatements := strings.Split(string(sqlBytes), ";")

	// Execute each statement
	for _, statement := range sqlStatements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		if err := db.Exec(statement).Error; err != nil {
			log.Printf("Error executing statement: %v\nStatement: %s", err, statement)
			return err
		}
	}

	return nil
}

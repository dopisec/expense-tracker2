package config

import (
	"expense-tracker-backend/models"
	"database/sql"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Import pure Go SQLite driver FIRST - this registers it as "sqlite" driver
// This MUST be imported before gorm.io/driver/sqlite tries to detect drivers
import _ "modernc.org/sqlite"

var DB *gorm.DB

func InitDB() error {
	// Create data directory if it doesn't exist
	dataDir := "data"
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return err
	}

	// Connect to SQLite database using pure Go driver explicitly
	// modernc.org/sqlite registers itself as "sqlite" driver name
	dbPath := filepath.Join(dataDir, "expense.db")
	
	// Open database connection using pure Go driver directly via database/sql
	// This ensures we use modernc.org/sqlite, not go-sqlite3
	sqlDB, err := sql.Open("sqlite", dbPath+"?_pragma=foreign_keys(1)")
	if err != nil {
		return err
	}
	
	// Test the connection to ensure driver is working
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return err
	}
	
	// Use Gorm's Open method with the existing sql.DB connection
	// This bypasses Gorm's driver detection and uses our explicit connection
	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		sqlDB.Close()
		return err
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Expense{})
	if err != nil {
		return err
	}

	return nil
}
